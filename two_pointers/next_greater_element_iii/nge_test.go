package next_greater_element_iii

import "testing"

func TestNge(t *testing.T) {
	t.Log(1<<31 - 1)
	r := nextGreaterElement(1223321)
	t.Log(r)
}
