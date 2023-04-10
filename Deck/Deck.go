package Deck

import (
	"errors"
	"math"
)

var invalidPriority = errors.New("invalid priority")
var emptyLowPriorityQueue = errors.New("low priority queue is empty")
var emptyHighPriorityQueue = errors.New("high priority queue is empty")

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

// Enqueue добавляет элемент в очеред в зависимости от приоритета (low или high)
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
	}

	// Добавляем элемент в очередь с высоким приоритетом
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

// DequeueLow забирает элемент с низким приоритетом
func (d *Deck) DequeueLow() (out int, err error) {
	// проверяем есть ли что забрать
	if d.LowPriorityCellCount == 0 {
		err = emptyLowPriorityQueue
		return
	}

	out = d.LowPriorityTopCell.next.value
	d.LowPriorityTopCell.next.next.prev = d.LowPriorityTopCell
	// обнуляем указатели у забираемой ячейки для предотвращения возможной утечки памяти
	d.LowPriorityTopCell.next.next, d.LowPriorityTopCell.next.prev, d.LowPriorityTopCell.next = nil, nil, d.LowPriorityTopCell.next.next

	// уменьшаем счетчик количества элементов в очереди с низким приоритетом
	d.LowPriorityCellCount--
	return
}

// DequeueHigh забирает элемент с высоким приоритетом
func (d *Deck) DequeueHigh() (out int, err error) {
	// проверяем есть ли что забрать
	if d.HighPriorityCellCount == 0 {
		err = emptyHighPriorityQueue
		return
	}

	out = d.HighPriorityTopCell.prev.value
	d.HighPriorityTopCell.prev.prev.next = d.HighPriorityTopCell
	// обнуляем указатели у забираемой ячейки для предотвращения возможной утечки памяти
	d.HighPriorityTopCell.prev.prev, d.HighPriorityTopCell.prev.next, d.HighPriorityTopCell.prev = nil, nil, d.HighPriorityTopCell.prev.prev

	// уменьшаем счетчик количества элементов в очереди с высоким приоритетом
	d.HighPriorityCellCount--
	return
}
