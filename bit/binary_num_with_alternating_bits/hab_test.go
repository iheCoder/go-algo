package binary_num_with_alternating_bits

import "testing"

func TestHAB(t *testing.T) {
	x := 1024
	y := x >> 1 & m1
	t.Log(y)
	z := x & m1 << 1
	t.Log(z)
	m := y | z
	t.Log(m)
	t.Log(m & x)
	x = 5
	t.Log(x >> 1)
	t.Log(x>>1 ^ x)
	t.Log(hasAlternatingBits(5))
}
