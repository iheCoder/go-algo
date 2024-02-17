package min_number_of_swaps_to_make_the_string_balanced

import "testing"

func TestMS(t *testing.T) {
	s := "][[]][][[][]"
	r := minSwaps(s)
	t.Log(r)
}
