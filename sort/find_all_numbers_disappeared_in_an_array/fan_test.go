package find_all_numbers_disappeared_in_an_array

import "testing"

func TestFAN(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 3}
	t.Log(findDisappearedNumbers(nums))
}
