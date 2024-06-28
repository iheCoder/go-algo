package custom_sort_str

import "testing"

func TestCSS(t *testing.T) {
	order := "bcafg"
	s := "abcd"
	t.Log(customSortString(order, s))
}
