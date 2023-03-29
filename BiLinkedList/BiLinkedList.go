package BiLinkedList

import (
	"errors"
	"fmt"
	"math"
)

type BiLinkedList struct {
	Top    *Cell
	Bottom *Cell
}

type Cell struct {
	Value int
	Next  *Cell
	Prev  *Cell
}

func insertCell(afterMe *Cell, value int) {
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

func (bl *BiLinkedList) AddToStart(value int) {
	insertCell(bl.Top, value)
	//newCell := &Cell{value, bl.Top.Next, bl.Top}
	//bl.Top.Next.Prev = newCell
	//bl.Top.Next = newCell
}

func (bl *BiLinkedList) AddToEnd(value int) {
	insertCell(bl.Bottom.Prev, value)
	//newCell := &Cell{value, bl.Bottom, bl.Bottom.Prev}
	//bl.Bottom.Prev.Next = newCell
	//bl.Bottom.Prev = newCell
}

func (bl *BiLinkedList) AddSort(value int) {
	next := bl.Top
	for next.Next.Value < value {
		next = next.Next
	}
	insertCell(next, value)
}

func (bl *BiLinkedList) Delete(value int) error {
	next := bl.Top.Next
	for next != bl.Bottom {
		if next.Value == value {
			next.Prev.Next, next.Next.Prev = next.Next, next.Prev
			return nil
		}
		next = next.Next
	}

	return errors.New("not found")
}

func (bl *BiLinkedList) Print() error {
	next := bl.Top.Next
	if next == bl.Bottom {
		return errors.New("list is empty")
	}
	for next != bl.Bottom {
		fmt.Print(next.Value, " ")
		next = next.Next
	}
	fmt.Println()
	return nil
}
