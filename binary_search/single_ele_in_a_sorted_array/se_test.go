package single_ele_in_a_sorted_array

import "testing"

func TestSE(t *testing.T) {
	nums := []int{3, 3, 7, 7, 10, 11, 11}
	t.Log(singleNonDuplicate(nums))
	nums = []int{1, 1, 2, 2, 4, 4, 5, 5, 9}
	t.Log(singleNonDuplicate(nums))
	nums = []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	t.Log(singleNonDuplicate(nums))
	nums = []int{1, 2, 2, 3, 3, 4, 4, 8, 8}
	t.Log(singleNonDuplicate(nums))
	nums = []int{3, 3, 7, 7, 11, 11, 12}
	t.Log(singleNonDuplicate(nums))
}
