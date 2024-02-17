package merge_intervals

import "testing"

func TestMI(t *testing.T) {
	is := [][]int{
		{1, 3},
		{2, 6},
		{6, 10},
		{1, 11},
	}
	t.Log(merge(is))
}
