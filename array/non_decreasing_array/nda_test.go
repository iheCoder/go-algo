package non_decreasing_array

import "testing"

func TestNDA(t *testing.T) {
	nums := []int{3, 4, 2, 3}
	t.Log(checkPossibility(nums))
}
