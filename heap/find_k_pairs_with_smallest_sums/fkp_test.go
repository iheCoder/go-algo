package find_k_pairs_with_smallest_sums

import "testing"

func TestKSP(t *testing.T) {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	t.Log(kSmallestPairs(nums1, nums2, 5))
}
