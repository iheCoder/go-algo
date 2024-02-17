package lexicographical_numbers

import "testing"

func TestLN(t *testing.T) {
	n := 12
	t.Log(lexicalOrder(n))
}
