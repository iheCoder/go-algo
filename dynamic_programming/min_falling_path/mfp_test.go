package min_falling_path

import "testing"

func TestMFP(t *testing.T) {
	matrix := [][]int{
		{-19, 57},
		{-40, -5},
	}
	t.Log(minFallingPathSum(matrix))
}
