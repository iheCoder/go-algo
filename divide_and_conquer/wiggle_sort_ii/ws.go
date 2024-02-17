package wiggle_sort_ii

import "sort"

// 采用quick sort规律
// 感觉合并的时候会很麻烦
func wiggleSort(nums []int) {
	if len(nums) == 1 {
		return
	}

	sort.Ints(nums)
	n := len(nums)
	r := make([]int, n)

	me := n >> 1
	if n&1 > 0 {
		r[n-1] = nums[0]
	} else {
		me--
	}

	for i := 0; i < n>>1; i++ {
		r[i<<1] = nums[me-i]
		r[i<<1+1] = nums[n-i-1]
	}

	for j, num := range r {
		nums[j] = num
	}
}

// 	n := len(nums)
//	n = n >> 1
//	c := n
//	if len(nums)&1 > 0 {
//		n++
//	}
//
//	var j, k int
//	for i := 0; i < c; i++ {
//		j, k = i<<1+1, n+i
//		nums[j], nums[k] = nums[k], nums[j]
//	}
