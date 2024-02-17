package set_mismatch

import "testing"

func TestSM(t *testing.T) {
	nums := []int{1, 2, 2, 4}
	t.Log(findErrorNums(nums))
	for i := 1; i < 10; i++ {
		t.Log(i & (-i))
	}
}
