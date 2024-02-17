package combination_sum

import "testing"

func TestCS(t *testing.T) {
	cs := []int{2, 3, 5}
	target := 8
	t.Log(combinationSum(cs, target))
}
