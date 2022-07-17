package order_by

import (
	"container/heap"
)

type OrderPriorityQueue []interface{}

//type OrderPriorityQueue struct {
//	orderByValues []interface{}
//}

func NewOrderPriorityQueue() OrderPriorityQueue {
	pq := OrderPriorityQueue{}
	heap.Init(&pq)
	return pq
}

func (pq *OrderPriorityQueue) Len() int {
	return len(*pq)
}

func (pq OrderPriorityQueue) Less(i, j int) bool {
	a := pq[i].(int)
	b := pq[j].(int)
	return a > b
}

func (pq OrderPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *OrderPriorityQueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
	//pq.update()
}

func (pq *OrderPriorityQueue) Pop() interface{} {
	return nil
}

func (pq *OrderPriorityQueue) Peek() interface{} {
	return nil
}

// update modifies the priority and value of an Item in the queue.
func (pq *OrderPriorityQueue) update() {
	heap.Fix(pq, len(*pq)-1)
}
