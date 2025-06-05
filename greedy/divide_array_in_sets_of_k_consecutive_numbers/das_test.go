package divide_array_in_sets_of_k_consecutive_numbers

import "testing"

func TestDAS(t *testing.T) {
	arr := []int{1, 1, 2, 2, 3, 3}
	t.Log(isPossibleDivide(arr, 2))
}
