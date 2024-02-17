package max_beauty_of_an_array_applying_op

import "sort"

func maximumBeauty(nums []int, k int) (ans int) {
	sort.Ints(nums)
	var left int
	for right, num := range nums {
		for num-nums[left] > 2*k {
			left++
		}
		ans = maxInt(ans, right-left+1)
	}
	return
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
