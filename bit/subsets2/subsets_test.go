package subsets2

import "testing"

func TestSubset2(t *testing.T) {
	nums := []int{1, 2, 2}
	t.Log(subsetsWithDup(nums))
}
