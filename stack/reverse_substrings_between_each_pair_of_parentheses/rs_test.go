package reverse_substrings_between_each_pair_of_parentheses

import "testing"

func TestRS(t *testing.T) {
	s := "(u(love)i)"
	t.Log(reverseParentheses(s))
}
