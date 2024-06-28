package num_of_subarrays_with_bounded_max

import "testing"

func TestSBM(t *testing.T) {
	nums := []int{2, 1, 4, 3}
	left := 2
	right := 3
	t.Log(numSubarrayBoundedMax(nums, left, right))
}
