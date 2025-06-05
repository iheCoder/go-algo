package count_num_of_nice_subarrays

import "testing"

func TestNS(t *testing.T) {
	nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	t.Log(numberOfSubarrays(nums, 2))
}
