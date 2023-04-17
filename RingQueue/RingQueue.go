package RingQueue

import "errors"

var capacityError = errors.New("queue is full")
var emptyQueueError = errors.New("queue is empty")

// RingQueue кольцевая очередь для элементов типа int
type RingQueue struct {
	values []any // содержит значения в очереди
	next   int   // указывает на следующий элемент для постановки в очередь
	last   int   // указывает на последний элемент в очереди
	length int   // тукущая длина кольца
}

// NewRingQueue создает новую очередь
func NewRingQueue(size int) *RingQueue {
	return &RingQueue{make([]any, size, size), 0, 0, 0}
}

// Enqueue добавляет элемент в очередь
func (rq *RingQueue) Enqueue(val any) (err error) {
	// если очередь полна, то возвращаем ошибку
	if rq.length == len(rq.values) {
		return capacityError
	}

	rq.values[rq.next] = val
	rq.next = (rq.next + 1) % len(rq.values)
	rq.length++
	return
}

// Dequeue забирает элемент из очереди
func (rq *RingQueue) Dequeue() (out any, err error) {
	// если очередь пустая, возвращаем ошибку
	if rq.length == 0 {
		err = emptyQueueError
		return
	}

	out = rq.values[rq.last]
	rq.values[rq.last] = 0 // для проверки функционала
	rq.last = (rq.last + 1) % len(rq.values)
	rq.length--
	return
}

// Top показывает следующий элемент в очереди
func (rq *RingQueue) Top() (top any, err error) {
	// если очередь пустая, возвращаем ошибку
	if rq.length == 0 {
		err = emptyQueueError
		return
	}

	top = rq.values[rq.last]
	return
}
