package complete_binary_tree_inserter

import "testing"

func TestPrintTestCases(t *testing.T) {
	n := 9900
	sc := ",\"insert\""
	var s string
	for i := 0; i < n; i++ {
		s += sc
	}
	t.Log(s)
	s = ""
	sc = ",[3]"
	for i := 0; i < n; i++ {
		s += sc
	}
	t.Log(s)
}
