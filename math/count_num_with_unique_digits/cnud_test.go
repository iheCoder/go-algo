package count_num_with_unique_digits

import "testing"

func TestCNUD(t *testing.T) {
	n := 4 // 5275
	t.Log(countNumbersWithUniqueDigits(n))
}

func TestCom(t *testing.T) {
	m := 4
	n := 2
	t.Log(combination(m, n))
}
