package min_flips_to_make_a_or_b_equal_to_c

import "testing"

func TestMF(t *testing.T) {
	a := 2
	b := 6
	c := 5
	t.Log(minFlips(a, b, c))
}
