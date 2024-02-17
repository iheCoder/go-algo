package number_of_islands

import "testing"

func TestNI(t *testing.T) {
	grid := [][]byte{
		{'1', '0', '1'},
		{'1', '1', '0'},
		{'0', '1', '0'},
	}
	//t.Log(numIslands(grid))
	grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	t.Log(numIslands(grid))

}
