package continous_subarray_sum

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	m := map[int]int{0: -1}
	var rem int
	for i, num := range nums {
		rem = (rem + num) % k
		if lastSameRemIndex, ok := m[rem]; ok {
			if i-lastSameRemIndex >= 2 {
				return true
			}
		} else {
			m[rem] = i
		}
	}

	return false
}
