package order_by

import (
	"testing"
)

func TestOrderPriorityQueue_Push(t *testing.T) {
	var pq OrderPriorityQueue = make([]interface{}, 0)
	pq.Push(2)
	pq.update()
	pq.Push(3)
	pq.update()
	pq.Push(1)
	pq.update()
	pq.Push(5)
	pq.update()
	pq.Push(2)
	pq.update()

	pq.Len()
}
