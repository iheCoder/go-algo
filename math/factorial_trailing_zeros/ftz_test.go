package factorial_trailing_zeros

import "testing"

func TestFactorial(t *testing.T) {
	for i := 1; i <= 15; i++ {
		t.Logf("%d \n %d", i, fac(i))
	}
}

func fac(n int) int {
	x := 1
	for i := 1; i <= n; i++ {
		x *= i
	}
	return x
}
