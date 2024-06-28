package max_area_of_island

import "testing"

func TestMAOI(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 0},
	}
	t.Log(maxAreaOfIsland(grid))
}
