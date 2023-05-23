package next_greater_element_iii

import (
	"math"
	"sort"
)

func nextGreaterElement(n int) int {
	arr := make([]int, 0)
	for n != 0 {
		arr = append(arr, n%10)
		n = n / 10
	}

	last := arr[0]
	var i int
	for i = 1; i < len(arr) && arr[i] >= last; i++ {
		last = arr[i]
	}
	if i >= len(arr) {
		return -1
	}

	j := sort.SearchInts(arr[:i], arr[i]+1)
	arr[i], arr[j] = arr[j], arr[i]

	i0, i1 := 0, i-1
	for i0 < i1 {
		arr[i0], arr[i1] = arr[i1], arr[i0]
		i0++
		i1--
	}

	var r int
	for k := len(arr) - 1; k >= 0; k-- {
		r += arr[k] * int(math.Pow10(k))
	}

	if r > 1<<31-1 {
		return -1
	}

	return r
}
