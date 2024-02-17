package h_index_ii

import (
	"sort"
	"testing"
)

func TestHIndex(t *testing.T) {
	cs := []int{100, 100, 100}
	t.Log(hIndex(cs))
	cs = []int{0}
	t.Log(hIndex(cs))
	cs = []int{0, 1, 2, 2, 2}
	t.Log(hIndex(cs))
	cs = []int{0, 1, 3, 5, 6, 7, 7, 7}
	t.Log(hIndex(cs))
	cs = []int{1, 1, 1, 1, 3, 3, 4, 4, 5, 6, 7, 7, 8, 9, 10}
	t.Log(hIndex(cs))
	cs = []int{2, 4, 7, 7, 7}
	t.Log(hIndex(cs))
}

func hIndex(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(i int) bool {
		// 如果引用数为x，所在位置为i，那么至少有n-i个引用数不小于x
		// 那么如果x不小于n-i，则意味着有可能取得更大的hIndex
		return citations[i] >= n-i
	})
}

//func hIndex(citations []int) int {
//	n := len(citations)
//	low, high := 0, n-1
//	var r int
//	for low < high {
//		x := low + (high-low)/2
//		r = maxInt(r, minInt(citations[x], n-x))
//		if x+1 == citations[x] {
//			return r
//		} else if x+1 < citations[x] || r < citations[x] {
//			high = x
//		} else {
//			low = x + 1
//		}
//	}
//	return r
//}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
