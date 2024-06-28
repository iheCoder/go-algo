package most_stones_removed_with_same_row_or_col

import "testing"

func TestMS(t *testing.T) {
	stones := [][]int{
		{1, 0},
		{0, 0},
		{0, 1},
		{1, 1},
	}
	t.Log(removeStones(stones))
}
