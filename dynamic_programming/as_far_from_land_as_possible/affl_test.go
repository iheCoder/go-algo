package as_far_from_land_as_possible

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenRandGrid(t *testing.T) {
	n := 100
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := rand.Intn(3)
			if x == 1 {
				grid[i][j] = 1
			}
		}
	}

	// Print the matrix in the desired format
	fmt.Print("[")
	for i, row := range grid {
		fmt.Print("[")
		for j, val := range row {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Print(val)
		}
		fmt.Print("]")
		if i < len(grid)-1 {
			fmt.Print(",\n ")
		}
	}
	fmt.Print("]\n")
}

func TestAffl(t *testing.T) {
	grid := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	t.Log(maxDistance(grid))
}
