package sqrtx

import (
	"math"
	"testing"
)

func TestSqrtx(t *testing.T) {
	x := 8
	//t.Log(mySqrt(x))
	x = 2
	t.Log(mySqrt(x))
	x = 0
	t.Log(mySqrt(x))
	for i := 0; i < 10000; i++ {
		expected := int(math.Sqrt(float64(i)))
		got := mySqrt(i)
		if expected != got {
			t.Fatalf("index %d got %d exp %d", i, got, expected)
		}
	}
	t.Log("PASS")
}
