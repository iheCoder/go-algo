package min_window_substring

import "testing"

func TestMWS(t *testing.T) {
	s := "ADOBECODEBANC"
	st := "ABC"
	//t.Log(minWindow(s, st))
	s = "ABCabc"
	st = "ab"
	//t.Log(minWindow(s, st))
	s = "ab"
	st = "a"
	//t.Log(minWindow(s, st))
	s = "abc"
	st = "cba"
	//t.Log(minWindow(s, st))
	s = "aa"
	st = "aa"
	t.Log(minWindow(s, st))
}
