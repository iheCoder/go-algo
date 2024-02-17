package fraction_to_recurring_decimal

import "testing"

func TestFRD(t *testing.T) {
	n, d := 4, 333
	//t.Log(fractionToDecimal(n, d))
	n = 1
	d = 90
	t.Log(fractionToDecimal(n, d))
}
