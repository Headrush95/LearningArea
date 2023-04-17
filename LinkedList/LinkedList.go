package LinkedList

import (
	"errors"
	"fmt"
	"math"
)

var valueNotFound = errors.New("value not found")
var emptyList = errors.New("list is empty")
var invalidType = errors.New("invalid type of element. Need integer")

// LinkedList связный список с верхним и нижним ограничителями
type LinkedList struct {
	Top    *Cell
	Bottom *Cell
}

// Cell ячейка связного списка
type Cell struct {
	Value any
	Next  *Cell
}

func NewLinkedList() *LinkedList {
	top := &Cell{math.MinInt64, nil}
	bottom := &Cell{math.MaxInt64, nil}
	bottom.Next = top // для добавления элементов в конец за O(1)
	return &LinkedList{Top: top, Bottom: bottom}
}

// AddToStart добавляет значение в начало списка
func (ll *LinkedList) AddToStart(value any) {
	newCell := &Cell{value, ll.Top.Next}
	ll.Top.Next = newCell
	if newCell.Next == nil {
		ll.Bottom.Next = newCell
	}
}

// AddToEnd добовляет значение в конец списка
func (ll *LinkedList) AddToEnd(value any) {
	newCell := &Cell{value, nil}
	ll.Bottom.Next.Next = newCell
	ll.Bottom.Next = newCell
}

// AddAfterMe вставляет значение после какой-то конретной ячейки. Если ее нет, то вставляет в конец
func (ll *LinkedList) AddAfterMe(afterMe any, value any) {
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

// checkType вспомогательная функция для AddSort и Max. Проверяет тип переменной
func checkType(val any) (err error) {
	if _, ok := val.(int); !ok {
		err = invalidType
		return
	}
	return
}

// AddSort добавляет значение по возрастанию. Имеет смысл только в сортированном списке
func (ll *LinkedList) AddSort(value int) (err error) {
	newCell := &Cell{value, nil}
	next := ll.Top
	for next.Next != nil {
		if err = checkType(next.Value); err != nil {
			return
		}
		if next.Next.Value.(int) >= value {
			break
		}
		next = next.Next
	}
	if next.Next == nil {
		ll.Bottom.Next = newCell
	}

	next.Next, newCell.Next = newCell, next.Next
	return
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

// Max возвращает максимальное значение в связном спике. ТОЛЬКО INTEGER
func (ll *LinkedList) Max() (max int, err error) {
	next := ll.Top.Next

	if err = checkType(ll.Top.Value); err != nil {
		return
	}
	max = ll.Top.Value.(int)

	for next != nil {
		if err = checkType(next.Value); err != nil {
			return
		}
		if next.Value.(int) > max {
			max = next.Value.(int)
		}
		next = next.Next
	}
	return
}

// IsSorted возвращет true, если элементы в связном списке отсортированны
func (ll *LinkedList) IsSorted() (sorted bool, err error) {
	next := ll.Top
	for next.Next != nil {
		if err = checkType(next.Next.Value); err != nil {
			return
		}
		if next.Value.(int) > next.Next.Value.(int) {
			return
		}
		next = next.Next
	}
	sorted = true
	return
}

// Delete удаляет элемент из списка, если такого нет - возвращает ошибку
func (ll *LinkedList) Delete(value any) (err error) {
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
