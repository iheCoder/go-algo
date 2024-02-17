package super_ugly_number

import "testing"

func TestSun(t *testing.T) {
	ps := []int{2, 7, 13, 19}
	n := 12
	t.Log(nthSuperUglyNumber(n, ps))
}
