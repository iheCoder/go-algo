package gen_parentheses

import "testing"

func TestGP(t *testing.T) {
	n := 3
	t.Log(generateParenthesis(n))
}
