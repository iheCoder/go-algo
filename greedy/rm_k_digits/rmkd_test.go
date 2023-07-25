package rm_k_digits

import "testing"

func TestRMKD(t *testing.T) {
	num := "1432219"
	k := 3
	//t.Log(removeKdigits(num, k))
	num = "10200"
	k = 1
	t.Log(removeKdigits(num, k))
}
