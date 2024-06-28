package max_sum_circular_subarray

func maxSubarraySumCircular(nums []int) int {
	ans := nums[0]
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] <= nums[i-1] {
			continue
		}
		ans = maxInt(ans, nums[i])
		if nums[i] <= 0 {
			continue
		}
		tmp := nums[i]
		step := 1
		j := i + 1
		for step < n {
			tmp += nums[j%n]
			if tmp <= 0 {
				break
			}
			ans = maxInt(ans, tmp)
			j++
			step++
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
