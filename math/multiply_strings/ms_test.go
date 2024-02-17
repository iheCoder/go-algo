package multiply_strings

import "testing"

func TestMul(t *testing.T) {
	num1, num2 := "123", "456"
	t.Log(multiply(num1, num2))
}
