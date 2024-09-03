package min_domino_rotations_for_equal_row

import "testing"

func TestMDR(t *testing.T) {
	tops := []int{2, 1, 2, 4, 2, 2}
	bottoms := []int{5, 2, 6, 2, 3, 2}
	t.Log(minDominoRotations(tops, bottoms))
}
