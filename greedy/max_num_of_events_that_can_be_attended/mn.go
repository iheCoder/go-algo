package max_num_of_events_that_can_be_attended

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})

	pq := &MinHeap{}
	heap.Init(pq)

	ans, i, n := 0, 0, len(events)

	for d := 1; d <= 100000; d++ {
		for pq.Len() > 0 && (*pq)[0] < d {
			heap.Pop(pq)
		}

		for i < n && events[i][0] == d {
			heap.Push(pq, events[i][1])
			i++
		}

		if pq.Len() > 0 {
			heap.Pop(pq)
			ans++
		}
	}

	return ans
}

// 小根堆实现（用于存储事件的结束时间）
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
