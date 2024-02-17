package wiggle_subsequence

import "testing"

type testData struct {
	nums []int
	exp  int
}

func TestWS(t *testing.T) {
	tds := []testData{
		{
			nums: []int{1, 7, 4, 9, 2, 5},
			exp:  6,
		},
		{
			nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			exp:  7,
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			exp:  2,
		},
		{
			nums: []int{1, 3, 2, 2, 1, 3},
			exp:  4,
		},
	}
	for i, td := range tds {
		got := wiggleMaxLength(td.nums)
		if got != td.exp {
			t.Fatalf("index %d exp %d got %d", i, td.exp, got)
		}
	}
}
