package jump_game_iii

import "testing"

func TestJG(t *testing.T) {
	arr, start := []int{4, 2, 3, 0, 3, 1, 2}, 0
	t.Log(canReach(arr, start))
}
