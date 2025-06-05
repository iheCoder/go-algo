package get_equal_substr_with_budget

import "testing"

func TestGesb(t *testing.T) {
	x, y := "abcd", "bcdf"
	//t.Log(equalSubstring(x, y, 3))
	x, y = "krrgw", "zjxss"
	t.Log(equalSubstring(x, y, 19))
}
