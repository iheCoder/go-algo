package h_index

import "sort"

// pair ? or check if size exceed citation
func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	var curIndex int
	for i := n - 1; i >= 0; i-- {
		if curIndex != citations[i] {
			if citations[i] > n-i-1 {
				curIndex = citations[i]
			} else {
				curIndex = n - i - 1
			}
		}
		if n-i >= curIndex {
			return curIndex
		}
	}
	return n
}
