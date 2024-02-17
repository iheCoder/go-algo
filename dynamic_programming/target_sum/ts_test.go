package target_sum

import "testing"

func TestTS(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	//t.Log(findTargetSumWays(nums, target))
	//nums = []int{1, 0}
	//target = 1
	//t.Log(findTargetSumWays(nums, target))
	nums = []int{0, 0}
	target = 0
	t.Log(findTargetSumWays(nums, target))
}
