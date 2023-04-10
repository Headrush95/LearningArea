package RabbitAndTurtleLoopCheck

import (
	"Projects/LinkedList"
	"fmt"
)

// rabbit указатель, который перемещается на 2 ячейки за шаг
type rabbit struct {
	step     int
	position *LinkedList.Cell
}

func (r *rabbit) move() {
	for i := r.step; i > 0; i-- {
		if r.position != nil {
			r.position = r.position.Next
		}
	}
}

// turtle указатель, который перемещается на 1 ячейку за шаг
type turtle struct {
	step     int
	position *LinkedList.Cell
}

func (t *turtle) move() {
	for i := t.step; i > 0; i-- {
		if t.position != nil {
			t.position = t.position.Next
		}
	}
}

// HasLoop проверяет однонаправленный связний список на наличие цикла
func HasLoop(cell *LinkedList.Cell) bool {
	r := rabbit{2, cell}
	t := turtle{1, cell}
	for r.position != nil {
		r.move()
		if r.position == nil {
			return false
		}
		t.move()
		if r.position == t.position {
			r.position = cell
			r.step = 1
			break
		}
	}

	for r.position != t.position {
		r.move()
		t.move()
	}
	r.step = 0
	for t.position.Next != r.position {
		t.move()
	}

	fmt.Printf("[HasLoop] loop starts on %d and ends on %d\n", r.position.Value, t.position.Value)

	return true
}
