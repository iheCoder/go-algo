package array_of_doubled_pairs

import "testing"

func TestCRD(t *testing.T) {
	arr := []int{1, 2, 1, -8, 8, -4, 4, -4, 2, -2}
	t.Log(canReorderDoubled(arr))
}
