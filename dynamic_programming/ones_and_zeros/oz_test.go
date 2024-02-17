package ones_and_zeros

import "testing"

func TestOZ(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	t.Log(findMaxForm(strs, m, n))
}
