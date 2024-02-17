package construct_quad_tree

import "testing"

func TestAlterGrid(t *testing.T) {
	grid := [][]int{
		{
			1,
			1,
			1,
			1,
		},
		{
			1,
			1,
			0,
			0,
		},
		{
			1,
			1,
			1,
			1,
		},
		{
			0,
			0,
			0,
			0,
		},
	}
	alterGrid(grid, 0, 0, len(grid))
	t.Log(grid)

	grid = [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0},
	}
	alterGrid(grid, 0, 0, len(grid))
	t.Log(grid)
}
