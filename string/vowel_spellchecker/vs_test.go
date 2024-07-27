package vowel_spellchecker

import "testing"

func TestSliceAppend(t *testing.T) {
	s := make([]int, 3)
	t.Log(s)
	_ = append(s, 1)
	t.Log(s)
	_ = append(s, 1)
	t.Log(s)
}
