package num_of_longest_increasing_subsequence

import "testing"

func TestNils(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7}
	t.Log(findNumberOfLIS(nums))
}
