package shortest_bridge

import "testing"

func TestSB(t *testing.T) {
	grid := [][]int{
		{0, 1},
		{1, 0},
	}

	//t.Log(shortestBridge(grid))

	grid = [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	//t.Log(shortestBridge(grid))

	grid = [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	//t.Log(shortestBridge(grid))

	//grid = [][]int{
	//	{1, 1, 0, 0, 0, 0},
	//	{0, 1, 0, 0, 0, 0},
	//	{1, 0, 0, 0, 0, 0},
	//	{1, 1, 1, 0, 0, 0},
	//	{1, 1, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0},
	//}
	//t.Log(shortestBridge(grid))

	grid = [][]int{
		{1, 0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	t.Log(shortestBridge(grid))
}
