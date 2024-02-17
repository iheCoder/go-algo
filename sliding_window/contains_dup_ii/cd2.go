package contains_dup_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	n := len(nums)
	m := make(map[int]bool)
	if k+1 > n {
		k = n - 1
	}
	for i := 0; i < k+1; i++ {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	i, j := 0, k+1
	for j < n {
		m[nums[i]] = false
		if m[nums[j]] {
			return true
		}
		m[nums[j]] = true
		i++
		j++
	}
	return false
}
