package jump_game1

import "testing"

func TestJG(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	//t.Log(jump(nums))
	nums = []int{2, 1}
	t.Log(jump(nums))
}
