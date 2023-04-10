package Stack

import (
	"errors"
)

var stackIsEmpty = errors.New("stack is empty")
var stackOverFlow = errors.New("stack overflow")

// Stack для элементов типа int
type Stack struct {
	values []int
	length int // длина
}

// NewStack возвращает Stack для типа int
func NewStack(size int) *Stack {
	return &Stack{make([]int, size, size), 0}
}

// Top показывает верхний элемент, но не забирает его
func (s *Stack) Top() (t int, err error) {
	if s.length == 0 {
		err = stackIsEmpty
		return
	}
	t = s.values[s.length-1]
	return
}

// Pop забирает и возвращает верхний элемент
func (s *Stack) Pop() (out int, err error) {
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
func (s *Stack) Push(val int) (err error) {
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

// InsertionSort сортирует элементы стека сортировкой вставкой
func InsertionSort(s *Stack, ascend bool) (err error) {
	if s.length == 0 {
		return stackIsEmpty
	}

	tempStack := NewStack(s.length)

	for i := 0; i < s.length; i++ {
		current, _ := s.Pop()

		// перемещаем неотсортированное
		for notSorted := 0; notSorted < s.length-1-i; notSorted++ {
			val, _ := s.Pop()
			_ = tempStack.Push(val)
		}

		// ищем в отсортированной части место для текущего элемента
		for s.length != 0 {
			val, _ := s.Pop()
			if ascend {
				if val >= current {
					_ = s.Push(val)
					break
				}
			} else {
				if val <= current {
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
		_ = tempStack.Push(val)
	}

	tmp := tempStack.length
	for i := 0; i < tmp; i++ {
		current, _ := tempStack.Pop()

		// ищем в отсортированной части место для текущего элемента
		for sorted := 0; sorted < i; sorted++ {
			val, _ := s.Pop()
			if ascend {
				if val >= current {
					current, val = val, current
				}
			} else {
				if val <= current {
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
		for notSorted := 0; notSorted < s.length-i; {
			val, _ := s.Pop()
			if ascend {
				if val >= current {
					val, current = current, val
				}
			} else {
				if val <= current {
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
