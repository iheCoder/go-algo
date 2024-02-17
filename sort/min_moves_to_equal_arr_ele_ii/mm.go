package min_moves_to_equal_arr_ele_ii

import "sort"

func minMoves2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)
	n := len(nums)
	var mid int
	if n&1 == 1 {
		mid = nums[n/2]
	} else {
		mid = (nums[n/2] + nums[n/2-1]) / 2
	}

	var ans int
	for _, num := range nums {
		ans += abs(num, mid)
	}
	return ans
}

// 平均值不行
// 	var sum int
//	for _, num := range nums {
//		sum += num
//	}
//	n := len(nums)
//	avg := sum / n
//
//	var ans int
//	for _, num := range nums {
//		ans += abs(num, avg)
//	}
//	return ans

func abs(x, y int) int {
	r := x - y
	if r < 0 {
		r = -r
	}
	return r
}
