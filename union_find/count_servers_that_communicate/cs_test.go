package count_servers_that_communicate

import "testing"

func TestCS(t *testing.T) {
	grid := [][]int{
		{1, 0},
		{0, 1},
	}
	t.Log(countServers(grid))
}
