package max_sum_with_exactly_k_elements

func maximizeSum(nums []int, k int) int {
	x := findMax(nums)
	return calSum(x, k)
}

func findMax(nums []int) int {
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > x {
			x = nums[i]
		}
	}
	return x
}

func calSum(m, k int) int {
	return (k * (2*m + k - 1)) / 2
}
