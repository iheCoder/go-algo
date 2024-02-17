package unique_substr_in_wraparound_str

import "testing"

func TestUS(t *testing.T) {
	s := "zab"
	t.Log(findSubstringInWraproundString(s))
}
