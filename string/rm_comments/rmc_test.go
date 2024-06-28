package rm_comments

import "testing"

func TestRMC(t *testing.T) {
	cs := []string{
		"a/*comment", "line", "more_comment*/b",
	}
	t.Log(removeComments(cs))
}
