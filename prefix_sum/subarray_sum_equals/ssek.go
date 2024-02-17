package subarray_sum_equals

func subarraySum(nums []int, k int) int {
	var sum int
	var ans int
	m := map[int]int{
		0: 1,
	}
	for _, num := range nums {
		sum += num
		if c, ok := m[sum-k]; ok {
			ans += c
		}
		m[sum]++
	}
	return ans
}
