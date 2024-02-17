package max_average_subarray_i

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	r := sum
	i, j := 0, k
	for j < len(nums) {
		sum -= nums[i]
		sum += nums[j]
		r = maxInt(r, sum)
		i++
		j++
	}
	return float64(r) / float64(k)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
