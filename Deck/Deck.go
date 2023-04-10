package Deck

import (
	"errors"
	"fmt"
	"math"
)

var invalidPriority = errors.New("invalid priority")
var emptyLowPriorityQueue = errors.New("low priority queue is empty")

// Deck очередь, в которой элементы с низким приоритетом ставятся в левую часть, а с высоким - в правую
type Deck struct {
	LowPriorityTopCell    *Cell
	LowPriorityCellCount  int
	HighPriorityTopCell   *Cell
	HighPriorityCellCount int
}

type Cell struct {
	value int
	next  *Cell
	prev  *Cell
}

// NewDeck возвращает новый дек
func NewDeck() *Deck {
	topLow := &Cell{math.MinInt64, nil, nil}
	topHigh := &Cell{math.MinInt64, nil, topLow}
	topLow.next = topHigh
	return &Deck{topLow, 0, topHigh, 0}
}

// Enqueue добавляет элемент в очеред в зависимости от приоритета
func (d *Deck) Enqueue(val int, priority string) (err error) {
	// проверка на валидность приоритета
	if priority != "low" && priority != "high" {
		err = invalidPriority
		return
	}
	if priority == "low" {
		low := d.LowPriorityTopCell
		// ищем последний элемент в очереди с низким приоритетом
		for i := 0; i < d.LowPriorityCellCount; i++ {
			low = low.next
		}
		newLowCell := &Cell{val, low.next, low}
		low.next.prev = newLowCell
		low.next = newLowCell
		d.LowPriorityCellCount++
		return
	} else if priority == "high" { // потом избавиться от else. Пока для проверки работы приоритета
		high := d.HighPriorityTopCell
		// ищем последний элемент в очереди с высоким приоритетом
		for i := 0; i < d.HighPriorityCellCount; i++ {
			high = high.prev
		}
		newHighCell := &Cell{val, high, high.prev}
		high.prev.next = newHighCell
		high.prev = newHighCell
		d.HighPriorityCellCount++
		return
	}
	fmt.Println("ошибка в логике")
	return
}

func (d *Deck) DequeueLow() (out int, err error) {
	// проверяем есть ли что забрать
	if d.LowPriorityCellCount == 0 {
		err = emptyLowPriorityQueue
		return
	}

	out = d.LowPriorityTopCell.next.value
	//d.LowPriorityTopCell.next.next.prev, d.LowPriorityTopCell.next.next = d.LowPriorityTopCell, nil
	//d.LowPriorityTopCell.next = d.LowPriorityTopCell.next.next

	return
}
