package find_1st_and_last_pos_of_ele_in_sorted_arr

import "testing"

func TestBS(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	t.Log(binarySearch(nums, target))
	target = 0
	t.Log(binarySearch(nums, target))
	target = 10
	t.Log(binarySearch(nums, target))
	target = 5
	t.Log(binarySearch(nums, target))
}

func TestSR(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	t.Log(searchRange(nums, target))
	target = 6
	t.Log(searchRange(nums, target))
	target = 5
	t.Log(searchRange(nums, target))
}
