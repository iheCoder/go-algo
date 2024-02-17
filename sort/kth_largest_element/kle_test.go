package kth_largest_element

import (
	"math"
	"testing"
)

func TestKLE(t *testing.T) {
	t.Log(math.MaxInt32)
	nums := []int{3, 2, 3, 1, 2, 4, 4, 5, 5, 6}
	t.Log(findKthLargest(nums, 4))
	nums = []int{1, 2, 2, 3, 2}
	t.Log(findKthLargest(nums, 2))
}
