package LinkedList

import (
	"errors"
	"fmt"
	"math"
)

// LinkedList связный список с верхним и нижним ограничителями
type LinkedList struct {
	Top    *Cell
	Bottom *Cell
}

// Cell ячейка связного списка
type Cell struct {
	Value int
	Next  *Cell
}

func NewLinkedList() *LinkedList {
	top := &Cell{math.MinInt64, nil}
	bottom := &Cell{math.MaxInt64, nil}
	top.Next = bottom
	bottom.Next = top // ALARM may cause infinite loop
	return &LinkedList{Top: top, Bottom: bottom}
}

// AddToStart добавляет значение в начало списка
func (ll *LinkedList) AddToStart(value int) {
	newCell := &Cell{value, ll.Top.Next}
	ll.Top.Next = newCell
}

// AddToEnd добовляет значение в конец списка
func (ll *LinkedList) AddToEnd(value int) {
	newCell := &Cell{value, ll.Bottom}
	ll.Bottom.Next.Next = newCell
	ll.Bottom.Next = newCell
}

// AddAfterMe вставляет значение после какой-то конретной ячейки. Если ее нет, то вставляет в конец
func (ll *LinkedList) AddAfterMe(afterMe int, value int) {
	newCell := &Cell{value, nil}
	next := ll.Top

	for next != ll.Bottom {
		if next.Value == afterMe {
			newCell.Next = next.Next
			next.Next = newCell
			return
		}
		next = next.Next
	}

	ll.AddToEnd(value)
}

// AddSort добавляет значение по возрастанию. Имеет смысл только в сортированном списке
func (ll *LinkedList) AddSort(value int) {
	newCell := &Cell{value, nil}
	next := ll.Top
	for next.Next.Value < value {
		next = next.Next
	}
	if next.Next == ll.Bottom {
		ll.Bottom.Next = newCell
	}
	newCell.Next = next.Next
	next.Next = newCell

}

// Print печатает список
func (ll *LinkedList) Print() error {
	next := ll.Top.Next
	if next == ll.Bottom {
		return errors.New("list is empty")
	}

	for next != ll.Bottom {
		fmt.Print(next.Value, " ")
		next = next.Next
	}

	fmt.Println()
	return nil
}

// Max возвращает максимальное значение в связном спике
func (ll *LinkedList) Max() int {
	next := ll.Top.Next
	max := ll.Top.Value
	for next != ll.Bottom {
		if next.Value > max {
			max = next.Value
		}
		next = next.Next
	}
	return max
}

// IsSorted возвращет true, если элементы в связном списке отсортированны
func (ll *LinkedList) IsSorted() bool {
	next := ll.Top
	for next != ll.Bottom {
		if next.Value > next.Next.Value {
			return false
		}
		next = next.Next
	}
	return true
}
