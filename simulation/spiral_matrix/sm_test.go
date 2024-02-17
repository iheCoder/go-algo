package spiral_matrix

import "testing"

func TestSM(t *testing.T) {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Log(spiralOrder(m))
}
