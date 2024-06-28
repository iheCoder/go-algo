package spiral_matrix_iii

import "testing"

func TestSM(t *testing.T) {
	r, c, rs, cs := 1, 4, 0, 0
	t.Log(spiralMatrixIII(r, c, rs, cs))
}
