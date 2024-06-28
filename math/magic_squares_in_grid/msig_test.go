package magic_squares_in_grid

import "testing"

func TestMsig(t *testing.T) {
	grid := [][]int{
		{
			4, 3, 8, 4,
		},
		{
			9, 5, 1, 9,
		},
		{
			2, 7, 6, 2,
		},
	}
	t.Log(numMagicSquaresInside(grid))
}
