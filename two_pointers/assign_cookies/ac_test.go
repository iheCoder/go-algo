package assign_cookies

import "testing"

type testData struct {
	g        []int
	s        []int
	expected int
}

func TestTPGreedy(t *testing.T) {
	tds := []testData{
		{
			[]int{10, 9, 8, 7},
			[]int{5, 6, 7, 8},
			2,
		},
	}

	for i, td := range tds {
		r := findContentChildren(td.g, td.s)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
