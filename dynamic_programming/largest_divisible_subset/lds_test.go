package largest_divisible_subset

import "testing"

func TestLDS(t *testing.T) {
	nums := []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}
	t.Log(largestDivisibleSubset(nums))
}
