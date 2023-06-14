package sum_of_two_int

import "testing"

func TestSub(t *testing.T) {
	a := 90
	b := 10
	t.Log((a ^ b) & (a | b>>1))
	a = 100
	b = 10
	t.Log((a ^ b) & (a | b>>1))
	t.Log((a ^ b) ^ ((a | ^b) << 1))
}

func TestSumOfTwoInt(t *testing.T) {
	a := 90
	b := -10
	t.Log(getSum(a, b))
}
