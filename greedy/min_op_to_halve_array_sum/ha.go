package min_op_to_halve_array_sum

import "container/heap"

type floatHeap []float64

func (h floatHeap) Len() int {
	return len(h)
}

func (h floatHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h floatHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *floatHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *floatHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func halveArray(nums []int) int {
	var sum float64
	for _, num := range nums {
		sum += float64(num)
	}

	target := sum / 2
	var fh floatHeap
	for _, num := range nums {
		fh = append(fh, float64(num))
	}
	heap.Init(&fh)
	var ans int
	var x float64
	for sum > target {
		x = (heap.Pop(&fh)).(float64)
		x = x / 2
		sum -= x
		heap.Push(&fh, x)
		ans++
	}
	return ans
}

// 	sort.Ints(nums)
//
//	n := len(nums)
//	arr := make([]float64, n)
//	for i, num := range nums {
//		arr[n-1-i] = float64(num)
//	}
// 		x := arr[0]
//		x /= 2
//		sum -= x
//		ans++
//		var i int
//		for i = 1; i < n && arr[i] > x; i++ {
//			arr[i-1] = arr[i]
//		}
//		arr[i-1] = x
