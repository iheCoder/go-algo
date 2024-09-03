package smallest_int_divisible_by_k

import "testing"

// 所有2倍数都不可能，以及5
// 其中3必定是7绑定，9和9绑定，1和1绑定
func TestEnsureOne(t *testing.T) {
	x := 111111
	y := x - 1
	for y > 1 {
		if x%y == 0 {
			t.Logf("found %d, mod %d", y, x/y)
		}
		y--
	}

	// 111 37, 3
	// 1111 101, 11
	// 11111 271, 41
}
