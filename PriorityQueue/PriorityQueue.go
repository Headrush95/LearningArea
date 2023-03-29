/*
PriorityQueue демонстрирует приоритетную очередь,
построенную с использованием интерфейса heap
*/
package PriorityQueue

import (
	"container/heap"
	"fmt"
	"math/rand"
)

// Item - это то, чем мы управляем в приоритетной очереди.
type Item struct {
	value    int // Значение элемента; произвольное.
	priority int // Приоритет элемента в очереди.
	// Индекс необходим для обновления
	// и поддерживается методами heap.Interface.
	index int // Индекс элемента в куче.
}

// PriorityQueue реализует heap.Interface и содержит Items.
type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// Мы хотим, чтобы Pop давал нам самый высокий,
	// а не самый низкий приоритет,
	// поэтому здесь мы используем оператор больше.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // избежать утечки памяти
	item.index = -1 // для безопасности
	*pq = old[0 : n-1]
	return item
}

// update изменяет приоритет и значение Item в очереди.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Этот пример создает PriorityQueue с некоторыми элементами,
// добавляет и управляет элементом,
// а затем удаляет элементы в порядке приоритета.
func main() {
	// Некоторые элементы и их приоритеты.
	items := make(map[int]int)
	var x int

	for i := 0; i < 10; i++ {
		x = rand.Intn(1000)
		items[x] = x
	}

	// Создаем очередь с приоритетами,
	// помещаем в нее элементы и
	// устанавливаем приоритетные инварианты очереди (кучи).
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Вставить новый элемент,
	// а затем изменить его приоритет.
	//item := &Item{
	//	value:    "orange",
	//	priority: 1,
	//}
	//heap.Push(&pq, item)
	//pq.update(item, item.value, 5)

	// Вынимаем предметы;
	// они прибывают в порядке убывания приоритета.
	//for pq.Len() > 0 {
	//	item := heap.Pop(&pq).(*Item)
	//	fmt.Printf("%.2d:%d ", item.priority, item.value)
	//}

	fmt.Println(heap.Pop(&pq).(*Item).value)
	fmt.Println(heap.Pop(&pq).(*Item).value)
	fmt.Println(heap.Pop(&pq).(*Item).value)
	fmt.Println(heap.Pop(&pq).(*Item).value)
}
