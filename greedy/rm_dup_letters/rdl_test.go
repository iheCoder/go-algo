package rm_dup_letters

import "testing"

func TestRDL(t *testing.T) {
	s := "cdadabcc"
	//t.Log(removeDuplicateLetters(s))
	//s = "abacb"
	//t.Log(removeDuplicateLetters(s))
	s = "bbcaac"
	t.Log(removeDuplicateLetters(s))
}
