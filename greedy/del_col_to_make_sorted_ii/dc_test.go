package del_col_to_make_sorted_ii

import "testing"

func TestDC(t *testing.T) {
	strs := []string{"ac", "ab", "aa"}
	t.Log(minDeletionSize(strs))
}
