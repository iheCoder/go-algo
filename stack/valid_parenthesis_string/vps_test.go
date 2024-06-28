package valid_parenthesis_string

import "testing"

func TestVPS(t *testing.T) {
	s := "(*)"
	t.Log(checkValidString(s))
	s = ")*)"
	t.Log(checkValidString(s))
	s = "*()(*)(*()()*((*)()))"
	t.Log(checkValidString(s))
	s = "((*()(*)(*()()*((*)()))"
	t.Log(checkValidString(s))
}
