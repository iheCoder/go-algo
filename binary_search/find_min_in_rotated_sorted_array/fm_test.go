package find_min_in_rotated_sorted_array

import "testing"

func TestFM(t *testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4}
	t.Log(findMin(nums))
	nums = []int{1, 2, 3}
	t.Log(findMin(nums))
	nums = []int{2, 3, 1}
	t.Log(findMin(nums))
	nums = []int{3, 1, 2}
	t.Log(findMin(nums))
	nums = []int{4, 1, 2, 3}
	t.Log(findMin(nums))
	nums = []int{2, 3, 4, 1}
	t.Log(findMin(nums))
	nums = []int{1, 2, 3, 4, 5}
	t.Log(findMin(nums))
	nums = []int{5, 1, 2, 3, 4}
	t.Log(findMin(nums))
}
