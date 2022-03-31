package heap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	queue := make(Queue, 10)
	for i := 0; i < 10; i++ {
		item := &Item{
			data:  i + 1,
			ref:   i + 1,
			index: i,
		}
		queue[i] = item
	}
	heap.Init(&queue)
	item := Item{
		data: 8,
		ref:  1,
	}
	heap.Push(&queue, &item)
	heap.Fix(&queue, 2)
	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		t.Log("index", item.index, "ref", item.ref, "val", item.data)
	}
}
