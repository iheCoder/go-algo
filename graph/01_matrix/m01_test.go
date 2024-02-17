package _1_matrix

import "testing"

func TestUM(t *testing.T) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
	}
	t.Log(updateMatrix(mat))
}
