package bag_of_tokens

import "testing"

type testData struct {
	tokens   []int
	power    int
	expected int
}

func TestBTGreedy(t *testing.T) {
	tds := []testData{
		{
			[]int{100},
			50,
			0,
		},
		{
			[]int{100, 200},
			150,
			1,
		},
		{
			[]int{100, 200, 300, 400},
			200,
			2,
		},
		{
			[]int{},
			200,
			0,
		},
		{
			[]int{100, 200, 300, 400},
			50,
			0,
		},
		{
			[]int{100, 200, 300, 400},
			100,
			1,
		},
		{
			[]int{6, 0, 39, 52, 45, 49, 59, 68, 42, 37},
			99,
			5,
		},
	}

	for i, td := range tds {
		r := bagOfTokensScore(td.tokens, td.power)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
