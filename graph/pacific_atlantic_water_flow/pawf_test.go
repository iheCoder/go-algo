package pacific_atlantic_water_flow

import "testing"

func TestPAWF(t *testing.T) {
	m := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	t.Log(pacificAtlantic(m))
}
