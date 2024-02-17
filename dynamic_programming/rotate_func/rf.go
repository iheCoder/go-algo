package rotate_func

func maxRotateFunction(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	n := len(nums)
	var sum, dsum int
	for i, num := range nums {
		sum += num
		dsum += i * num
	}

	ans := dsum
	for i := n - 1; i >= 0; i-- {
		dsum += sum - nums[i] - (n-1)*nums[i]
		if dsum > ans {
			ans = dsum
		}
	}
	return ans
}
