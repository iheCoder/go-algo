package find_and_replace_in_string

import "testing"

func TestFrs(t *testing.T) {
	s := "abcd"
	is := []int{0, 2}
	ss := []string{"a", "cd"}
	ts := []string{"eee", "ffff"}
	//t.Log(findReplaceString(s, is, ss, ts))
	s = "abcde"
	is = []int{2, 2}
	ss = []string{"bc", "cde"}
	ts = []string{"fe", "f"}
	t.Log(findReplaceString(s, is, ss, ts))
}
