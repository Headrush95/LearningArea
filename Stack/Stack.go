package Stack

import (
	"errors"
)

var stackIsEmpty = errors.New("stack is empty")
var stackOverFlow = errors.New("stack overflow")
var invalidType = errors.New("invalid type of element. Need integer")

// Stack для элементов любого типа
type Stack struct {
	values []any
	length int // длина
}

// NewStack возвращает Stack для типа int
func NewStack(size int) *Stack {
	return &Stack{make([]any, size, size), 0}
}

// Top показывает верхний элемент, но не забирает его
func (s *Stack) Top() (top any, err error) {
	if s.length == 0 {
		err = stackIsEmpty
		return
	}
	top = s.values[s.length-1]
	return
}

// Pop забирает и возвращает верхний элемент
func (s *Stack) Pop() (out any, err error) {
	if s.length == 0 {
		err = stackIsEmpty
		return
	}
	out = s.values[s.length-1]
	s.values[s.length-1] = 0 // для удобства проверки других функций
	s.length--
	return
}

// Push добавляет элемент в стек
func (s *Stack) Push(val any) (err error) {
	if s.length == len(s.values) {
		err = stackOverFlow
		return
	}
	s.values[s.length] = val
	s.length++
	return
}

// Len возвращает длину стека
func (s *Stack) Len() int {
	return s.length
}

// checkType вспомогательная функция для сортировок. Проверяет тип переменной
func checkType(val any) (err error) {
	if _, ok := val.(int); !ok {
		err = invalidType
		return
	}
	return
}

// InsertionSort сортирует элементы стека сортировкой вставкой
func InsertionSort(s *Stack, ascend bool) (err error) {
	if s.length == 0 {
		return stackIsEmpty
	}

	tempStack := NewStack(s.length)

	for i := 0; i < s.length; i++ {
		current, _ := s.Pop()
		if err = checkType(current); err != nil {
			err = invalidType
			return
		}

		// перемещаем неотсортированное
		for notSorted := 0; notSorted < s.length-1-i; notSorted++ {
			val, _ := s.Pop()
			if err = checkType(val); err != nil {
				return
			}

			_ = tempStack.Push(val)
		}

		// ищем в отсортированной части место для текущего элемента
		for s.length != 0 {
			val, _ := s.Pop()
			if ascend {
				if val.(int) >= current.(int) {
					_ = s.Push(val)
					break
				}
			} else {
				if val.(int) <= current.(int) {
					_ = s.Push(val)
					break
				}
			}

			_ = tempStack.Push(val)
		}
		_ = s.Push(current)

		// переносим все неотсортированное из временного стека в основной
		for tempStack.length != 0 {
			val, _ := tempStack.Pop()
			_ = s.Push(val)
		}

	}
	return
}

// InsertionSortMod делает тоже самое, что и InsertionSort, но все неотсортированные элементы хранятся во временном стеке, а отсортированные в основном
func InsertionSortMod(s *Stack, ascend bool) (err error) {
	if s.length == 0 {
		return stackIsEmpty
	}

	tempStack := NewStack(s.length)

	// перемещаем неотсортированное
	for s.length != 0 {
		val, _ := s.Pop()
		if err = checkType(val); err != nil {
			return
		}

		_ = tempStack.Push(val)
	}

	tmp := tempStack.length
	for i := 0; i < tmp; i++ {
		current, _ := tempStack.Pop()

		// ищем в отсортированной части место для текущего элемента
		for sorted := 0; sorted < i; sorted++ {
			val, _ := s.Pop()
			if ascend {
				if val.(int) >= current.(int) {
					current, val = val, current
				}
			} else {
				if val.(int) <= current.(int) {
					current, val = val, current
				}
			}
			_ = tempStack.Push(val)

		}

		// помещаем текущий элемент на свое место
		_ = s.Push(current)

		// возвращаем все отсортированное из временного стека в основной
		for back := 0; back < i; back++ {
			val, _ := tempStack.Pop()
			_ = s.Push(val)
		}

	}
	return
}

// SelectionSort сортирует элементы стека сортировкой выбором
func SelectionSort(s *Stack, ascend bool) (err error) {
	if s.length == 0 {
		return stackIsEmpty
	}

	tempStack := NewStack(s.length)

	// ищем самый меньший/больший элемент
	for i := 0; i < s.length; i++ {
		current, _ := s.Pop()
		if err = checkType(current); err != nil {
			return
		}

		for notSorted := 0; notSorted < s.length-i; {
			val, _ := s.Pop()
			if err = checkType(val); err != nil {
				return
			}

			if ascend {
				if val.(int) >= current.(int) {
					val, current = current, val
				}
			} else {
				if val.(int) <= current.(int) {
					val, current = current, val
				}
			}
			_ = tempStack.Push(val)
		}
		_ = s.Push(current)

		// возвращаем все элементы из temp в основной
		for tempStack.length != 0 {
			val, _ := tempStack.Pop()
			_ = s.Push(val)
		}
	}
	return
}
