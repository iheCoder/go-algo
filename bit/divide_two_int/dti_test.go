package divide_two_int

import "testing"

func TestDivideTwoInt(t *testing.T) {
	x := 4
	y := 2
	t.Log(divideAdd(x, y) == x/y) // 2
	x = 7
	y = -3
	t.Log(divideAdd(x, y) == x/y) // -2
	x = -2147483648
	y = 1
	t.Log(divideAdd(x, y) == x/y)
	y = -2
	t.Log(divideAdd(x, y) == x/y)
	x = -4
	t.Log(divideAdd(x, y) == x/y)
	x = 1
	y = 3
	t.Log(divideAdd(x, y) == x/y)
	x = 0
	t.Log(divideAdd(x, y) == x/y) // 0
}
