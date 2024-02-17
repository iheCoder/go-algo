package complex_number_multiplication

import "testing"

func TestCNM(t *testing.T) {
	num1 := "1+-1i"
	num2 := "1+1i"
	t.Log(complexNumberMultiply(num1, num2))
}
