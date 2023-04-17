package Queue

import (
	"errors"
	"fmt"
)

var emptyQueue = errors.New("queue is empty")
var invalidType = errors.New("invalid type of element. Need integer")

// Queue очередь в виде связного списка
type Queue struct {
	top      *Cell
	bottom   *Cell
	quantity int
}

// Cell ячейка связного списка. Тип элентов любой
type Cell struct {
	value    any
	nextCell *Cell
	prevCell *Cell
}

// NewQueue создает и возвращает новую пустую очередь
func NewQueue() *Queue {
	top := &Cell{nil, nil, nil}
	bottom := &Cell{nil, nil, top}
	top.nextCell = bottom
	return &Queue{top, bottom, 0}
}

// Enqueue помещает элемент в очередь
func (q *Queue) Enqueue(value any) {
	newCell := &Cell{value, q.top.nextCell, q.top}
	q.top.nextCell.prevCell = newCell
	q.top.nextCell = newCell

	q.quantity++
}

// Dequeue забирает элемент из очереди
func (q *Queue) Dequeue() (out any, err error) {
	//проверяем, есть ли элементы в очереди
	if q.top.nextCell == q.bottom {
		err = emptyQueue
		return
	}

	// на всякий случай еще обнуляем указатели забираемого элемента
	out = q.bottom.prevCell.value
	q.bottom.prevCell.nextCell = q.bottom
	q.bottom.prevCell.nextCell = nil
	q.bottom.prevCell.prevCell, q.bottom.prevCell = nil, q.bottom.prevCell.prevCell
	q.bottom.prevCell.nextCell = q.bottom
	q.quantity--

	return
}

// Len возвращает длину очереди
func (q *Queue) Len() (length int) {
	return q.quantity
}

// Print печататет все элементы в очереди начиная с первого
func (q *Queue) Print() (err error) {
	//проверяем, есть ли элементы в очереди
	if q.top.nextCell == q.bottom {
		err = emptyQueue
		return
	}
	next := q.bottom.prevCell
	for i := 0; i < q.quantity; i++ {
		fmt.Print(next.value, " ")
		next = next.prevCell
	}
	fmt.Println()
	return
}

// sortFix возращает на первое место в очереди элемент, не прошедший проверку в сортировках
func (q *Queue) sortFix(val any) {
	newCell := &Cell{val, q.bottom, q.bottom.prevCell}
	q.bottom.prevCell.nextCell = newCell
	q.bottom.prevCell = newCell
	q.quantity++
}

// InsertionSort сортирует элементы очереди сортировкой вставкой. ТОЛЬКО ДЛЯ ЭЛЕМЕНТОВ ТИПА INT
func (q *Queue) InsertionSort() (err error) {
	// проверяем не пуста ли очередь
	if q.top.nextCell == q.bottom {
		err = emptyQueue
		return
	}
	tempQueue := NewQueue()

	length := q.quantity
	for sorted := 0; sorted < length; sorted++ {
		current, _ := q.Dequeue()
		if _, ok := current.(int); !ok {
			err = invalidType
			q.sortFix(current)
			return
		}

		// перемещаем все неотсортированные элементы во временную очередь
		for notSorted := 0; notSorted < length-1-sorted; notSorted++ {
			val, _ := q.Dequeue()
			if _, ok := val.(int); !ok {
				err = invalidType
				return
			}
			tempQueue.Enqueue(val)
		}

		// ищем в отсортированной части место для текущего элемента
		for q.top.nextCell != q.bottom {
			val, _ := q.Dequeue()

			if val.(int) <= current.(int) {
				q.Enqueue(val)
				break
			}
			tempQueue.Enqueue(val)

		}
		q.Enqueue(current)

		// возвращаем элементы из временной очереди в основную
		for tempQueue.top.nextCell != tempQueue.bottom {
			val, _ := tempQueue.Dequeue()
			q.Enqueue(val)
		}
	}

	return
}

// SelectionSort сортирует элементы очереди сортировкой выбором. ТОЛЬКО ДЛЯ ЭЛЕМЕНТОВ ТИПА INT
func (q *Queue) SelectionSort() (err error) {
	// проверяем не пуста ли очередь
	if q.top.nextCell == q.bottom {
		err = emptyQueue
		return
	}
	tempQueue := NewQueue()

	length := q.quantity
	for sorted := 0; sorted < length; sorted++ {
		min, _ := q.Dequeue()

		//перемещаем все неотсортированные элементы во временный массив и отслеживаем минимальное значение
		for notSorted := length - sorted - 1; notSorted > 0; notSorted-- {
			curr, _ := q.Dequeue()
			if curr.(int) < min.(int) {
				curr, min = min, curr
			}
			tempQueue.Enqueue(curr)
		}

		// перекидываем отсортированную часть в конец временного массива
		for q.top.nextCell != q.bottom {
			tmp, _ := q.Dequeue()
			tempQueue.Enqueue(tmp)
		}
		// возвращаем все из временного массива в основной
		for tempQueue.top.nextCell != tempQueue.bottom {
			tmp, _ := tempQueue.Dequeue()
			q.Enqueue(tmp)
		}
		q.Enqueue(min)
	}

	return
}
