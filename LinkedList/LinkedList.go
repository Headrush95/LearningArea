package LinkedList

import (
	"errors"
	"fmt"
	"math"
)

var valueNotFound = errors.New("value not found")
var emptyList = errors.New("list is empty")

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
	bottom.Next = top // для добавления элементов в конец за O(1)
	return &LinkedList{Top: top, Bottom: bottom}
}

// AddToStart добавляет значение в начало списка
func (ll *LinkedList) AddToStart(value int) {
	newCell := &Cell{value, ll.Top.Next}
	ll.Top.Next = newCell
	if newCell.Next == nil {
		ll.Bottom.Next = newCell
	}
}

// AddToEnd добовляет значение в конец списка
func (ll *LinkedList) AddToEnd(value int) {
	newCell := &Cell{value, nil}
	ll.Bottom.Next.Next = newCell
	ll.Bottom.Next = newCell
}

// AddAfterMe вставляет значение после какой-то конретной ячейки. Если ее нет, то вставляет в конец
func (ll *LinkedList) AddAfterMe(afterMe int, value int) {
	newCell := &Cell{value, nil}
	next := ll.Top

	for next != nil {
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
	for next.Next != nil {
		if next.Next.Value >= value {
			break
		}
		next = next.Next
	}
	if next.Next == nil {
		ll.Bottom.Next = newCell
	}

	next.Next, newCell.Next = newCell, next.Next

}

// Print печатает список
func (ll *LinkedList) Print() (err error) {
	next := ll.Top.Next
	if next == nil {
		err = emptyList
		return
	}

	for next != nil {
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
	for next != nil {
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
	for next.Next != nil {
		if next.Value > next.Next.Value {
			return false
		}
		next = next.Next
	}
	return true
}

// Delete удаляет элемент из списка, если такого нет - возвращает ошибку
func (ll *LinkedList) Delete(value int) (err error) {
	next := ll.Top

	for next.Next != nil {
		// ищем нужный элемент
		if next.Next.Value == value {
			break
		}
		next = next.Next
	}

	// если дошли до конца и не нашли возвращаем ошибку
	if next.Next == nil {
		err = valueNotFound
		return
	}

	next.Next, next.Next.Next = next.Next.Next, nil

	// если удалили последний элемент, то обновляем поле Bottom.Next
	if next.Next == nil {
		ll.Bottom.Next = next
	}
	return
}
