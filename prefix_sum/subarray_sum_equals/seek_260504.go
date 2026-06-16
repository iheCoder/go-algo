package subarray_sum_equals

func subarraySum260504(nums []int, k int) int {
	m := map[int]int{
		0: 1,
	}

	var pre, ans int
	for _, num := range nums {
		pre += num
		ans += m[pre-k]
		m[pre]++
	}

	return ans
}
