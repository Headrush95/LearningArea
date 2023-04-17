package BiLinkedList

import (
	"errors"
	"fmt"
	"math"
)

var notFound = errors.New("value not found")
var emptyList = errors.New("list is empty")
var invalidType = errors.New("invalid type of element. Need integer")

// BiLinkedList двусвязный список с верхним и нижним ограничителем
type BiLinkedList struct {
	Top    *Cell
	Bottom *Cell
}

// Cell ячейка двусвязного списка
type Cell struct {
	Value any
	Next  *Cell
	Prev  *Cell
}

// insertCell приватная функция, которая вставляет элемент после afterMe
func insertCell(afterMe *Cell, value any) {
	newCell := &Cell{value, afterMe.Next, afterMe}
	afterMe.Next.Prev = newCell
	afterMe.Next = newCell
}

func NewBiLinkedList() *BiLinkedList {
	top := &Cell{math.MinInt64, nil, nil}
	bottom := &Cell{math.MaxInt64, nil, nil}
	top.Next = bottom
	bottom.Prev = top
	return &BiLinkedList{top, bottom}
}

// AddToStart добавляет элемент в начало двусвязного списка
func (bl *BiLinkedList) AddToStart(value any) {
	insertCell(bl.Top, value)

}

// AddToEnd добавляет элемент в конец двусвязного списка
func (bl *BiLinkedList) AddToEnd(value any) {
	insertCell(bl.Bottom.Prev, value)
}

// checkType вспомогательная функция для AddSort. Проверяет тип переменной
func checkType(val any) (err error) {
	if _, ok := val.(int); !ok {
		err = invalidType
		return
	}
	return
}

// AddSort добавляет элемент в сортированный список. ТОЛЬКО INTEGER
func (bl *BiLinkedList) AddSort(value int) (err error) {
	next := bl.Top
	if err = checkType(next.Next.Value); err != nil {
		return
	}
	for next.Next.Value.(int) < value {
		next = next.Next
		if err = checkType(next.Next.Value); err != nil {
			return
		}
	}
	insertCell(next, value)
	return
}

// DeleteCell удаляет ячейку по известному адресу
func DeleteCell(c *Cell) {
	c.Prev.Next, c.Next.Prev = c.Next, c.Prev
	// обнуляем ссылки для предотвращения утечек
	c.Next, c.Prev = nil, nil
	return
}

// Delete удаляет элемент по значению. Предварительно его ищет -> асимптотика O(n)
func (bl *BiLinkedList) Delete(value any) (err error) {
	next := bl.Top.Next
	for next != bl.Bottom {
		if next.Value == value {
			DeleteCell(next)
			return
		}
		next = next.Next
	}
	err = notFound
	return
}

// Print печатает весь список
func (bl *BiLinkedList) Print() (err error) {
	next := bl.Top.Next
	if next == bl.Bottom {
		err = emptyList
		return
	}
	for next != bl.Bottom {
		fmt.Print(next.Value, " ")
		next = next.Next
	}
	fmt.Println()
	return nil
}
