package find_right_interval

import "testing"

func TestFRI(t *testing.T) {
	is := [][]int{
		{3, 4},
		{2, 3},
		{1, 2},
	}
	t.Log(findRightInterval(is))
}
