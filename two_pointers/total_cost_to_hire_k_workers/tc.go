package total_cost_to_hire_k_workers

import (
	"container/heap"
	"math/rand"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h *intHeap) getMin() int {
	if h.Len() <= 0 {
		return 1<<31 - 1
	}
	return (*h)[0]
}

func totalCost(costs []int, k int, candidates int) int64 {
	if k >= len(costs) {
		var sum int
		for i := 0; i < len(costs); i++ {
			sum += costs[i]
		}
		return int64(sum)
	}
	if candidates >= len(costs) {
		return int64(randomizedSelectSum(costs, k))
	}

	var hl intHeap
	var li int
	for li = 0; li < candidates; li++ {
		hl = append(hl, costs[li])
	}

	var hr intHeap
	var ri int
	// 应该能等于，可能恰好就是那个最小数
	for ri = len(costs) - 1; ri >= len(costs)-candidates && ri >= li; ri-- {
		hr = append(hr, costs[ri])
	}

	heap.Init(&hl)
	heap.Init(&hr)

	var left, right int
	var sum int
	for i := 0; i < k; i++ {
		if hl.Len() == 0 && hr.Len() == 0 {
			return int64(sum)
		}

		left, right = hl.getMin(), hr.getMin()
		if left <= right {
			sum += left
			heap.Pop(&hl)
			if li <= ri {
				heap.Push(&hl, costs[li])
				li++
			}
		} else {
			sum += right
			heap.Pop(&hr)
			// TODO 为什么ri>=li就对了呢？
			// 不是ri>=li就对了，而是在li、ri相同的时候，恰好选择了left，right，那么就必须作出去选择li或者ri
			if ri >= li {
				heap.Push(&hr, costs[ri])
				ri--
			}
		}
	}

	return int64(sum)
}

func randomizedSelectSum(a []int, k int) int {
	if len(a) == 1 {
		return a[0]
	}

	pivot := rand.Intn(len(a))
	pivot = partition(a, pivot)

	if k < pivot {
		randomizedSelectSum(a[:pivot], k)
	} else if k > pivot {
		randomizedSelectSum(a[pivot+1:], k-pivot-1)
	}

	var sum int
	for i := 0; i < k; i++ {
		sum += a[i]
	}
	return sum
}

func partition(a []int, pivot int) int {
	a[pivot], a[len(a)-1] = a[len(a)-1], a[pivot]

	var j int
	for i := 0; i < len(a); i++ {
		if a[i] < a[len(a)-1] {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}

	a[j], a[len(a)-1] = a[len(a)-1], a[j]

	return j
}
