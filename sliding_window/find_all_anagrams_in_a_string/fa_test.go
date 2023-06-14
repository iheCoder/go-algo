package find_all_anagrams_in_a_string

import "testing"

func TestFA(t *testing.T) {
	s := "abab"
	p := "ab"
	//t.Log(findAnagrams(s, p))
	s = "cbaebabacd"
	p = "abc"
	t.Log(findAnagrams(s, p))
}
