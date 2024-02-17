package arithmetic_slices

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	d, t := nums[0]-nums[1], 0
	var ans int
	for i := 2; i < n; i++ {
		if d == nums[i-1]-nums[i] {
			t++
		} else {
			t = 0
		}
		d = nums[i-1] - nums[i]
		ans += t
	}
	return ans
}
