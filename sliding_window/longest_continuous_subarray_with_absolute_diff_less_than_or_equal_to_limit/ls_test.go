package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

import "testing"

func TestLS(t *testing.T) {
	nums := []int{8, 2, 4, 7}
	limit := 4
	//t.Log(longestSubarray(nums, limit))
	nums = []int{10, 1, 2, 4, 7, 2}
	limit = 5
	//t.Log(longestSubarray(nums, limit))
	nums = []int{4, 8, 5, 1, 7, 9}
	limit = 6
	t.Log(longestSubarray(nums, limit))
}
