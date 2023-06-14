package max_length_of_repeated_subarray

import "testing"

func TestFL(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	t.Log(findLength(nums1, nums2))
}
