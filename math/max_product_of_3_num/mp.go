package max_product_of_3_num

import "sort"

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)
	n := len(nums)
	negMax := nums[0] * nums[1] * nums[n-1]
	posMax := nums[n-2] * nums[n-3] * nums[n-1]
	if negMax < posMax {
		return posMax
	}
	return negMax
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
