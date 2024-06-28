package valid_square

import "testing"

func TestVS(t *testing.T) {
	p1 := []int{0, 1}
	p2 := []int{1, 2}
	p3 := []int{0, 2}
	p4 := []int{0, 0}

	p1 = []int{1, 0}
	p2 = []int{0, 1}
	p3 = []int{-1, 0}
	p4 = []int{0, -1}
	t.Log(validSquare(p1, p2, p3, p4))
}
