package zigzeg_conv

import "testing"

func TestZC(t *testing.T) {
	s := "PAYPALISHIRING"
	r := "PAHNAPLSIIGYIR"
	n := 3
	t.Log(convert(s, n) == r)
	s = "AB"
	n = 1
	t.Log(convert(s, n))
}
