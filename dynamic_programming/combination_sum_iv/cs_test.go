package combination_sum_iv

import "testing"

func TestCS(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 4
	//t.Log(combinationSum4Dep(nums, target))

	target = 5
	t.Log(combinationSum4Dep(nums, target))
}
