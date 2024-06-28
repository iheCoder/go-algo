package loud_and_rich

import "testing"

func TestLR(t *testing.T) {
	richer := [][]int{
		{1, 0},
		{2, 1},
		{3, 1},
		{3, 7},
		{4, 3},
		{5, 3},
		{6, 3},
	}
	quiet := []int{3, 2, 5, 4, 6, 1, 7, 0}
	t.Log(loudAndRich(richer, quiet))
}
