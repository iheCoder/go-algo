package int_replace

import "testing"

func TestIR(t *testing.T) {
	t.Log(10 & -10)
	t.Log(16&-16 > 14&-14)
	x := 999
	t.Log(integerReplacement(x))
}
