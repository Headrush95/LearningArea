package DoubleStack

import "errors"

var doubleStackIsFull = errors.New("double stack is full")
var leftPartEmpty = errors.New("left part of double stack is empty")
var rightPartEmpty = errors.New("right part of double stack is empty")

// DoubleStack двойная очередь для любых элементов. Элементы добавляются с двух концов
type DoubleStack struct {
	values     []any
	indexLeft  int
	indexRight int
}

// NewDoubleStack создает новый двойной стек для элементов типа int
func NewDoubleStack(size int) *DoubleStack {
	values := make([]any, size, size)
	return &DoubleStack{values, 0, len(values) - 1}
}

// PushLeft добавляет элемент в левую часть очереди
func (d *DoubleStack) PushLeft(val any) (err error) {
	// проверка на переполнение очереди
	if d.indexLeft > d.indexRight {
		err = doubleStackIsFull
		return
	}

	d.values[d.indexLeft] = val
	d.indexLeft++
	return
}

// PushRight добавляет элемент в правую часть очереди
func (d *DoubleStack) PushRight(val any) (err error) {
	// проверка на переполнение очереди
	if d.indexLeft > d.indexRight {
		err = doubleStackIsFull
		return
	}

	d.values[d.indexRight] = val
	d.indexRight--
	return
}

// PopLeft забирает элемент из левой части очереди
func (d *DoubleStack) PopLeft() (out any, err error) {
	// проверка на наличие элементов в левой части очереди
	if d.indexLeft == 0 {
		err = leftPartEmpty
		return
	}

	out = d.values[d.indexLeft-1]
	d.values[d.indexLeft-1] = 0 // для проверки
	d.indexLeft--
	return
}

// PopRight забирает элемент из правой части очереди
func (d *DoubleStack) PopRight() (out any, err error) {
	// проверка на наличие элементов в правой части очереди
	if d.indexRight == len(d.values)-1 {
		err = rightPartEmpty
		return
	}

	out = d.values[d.indexRight+1]
	d.values[d.indexRight+1] = 0 // для проверки
	d.indexRight++
	return
}

// TopLeft показывает верхний элемент слева
func (d *DoubleStack) TopLeft() (top any, err error) {
	// проверка на наличие элементов в левой части очереди
	if d.indexLeft == 0 {
		err = leftPartEmpty
		return
	}

	top = d.values[d.indexLeft-1]
	return
}

// TopRight показывает верхний элемент справа
func (d *DoubleStack) TopRight() (top any, err error) {
	// проверка на наличие элементов в правой части очереди
	if d.indexRight == len(d.values)-1 {
		err = rightPartEmpty
		return
	}

	top = d.values[d.indexRight+1]
	return
}
