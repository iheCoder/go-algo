package magical_string

import "testing"

type testData struct {
	n        int
	expected int
}

func TestMS(t *testing.T) {
	tds := []testData{
		{
			n:        6,
			expected: 3,
		},
		{
			n:        5,
			expected: 3,
		},
		{
			n:        7,
			expected: 4,
		},
		{
			n:        8,
			expected: 4,
		},
		{
			n:        9,
			expected: 4,
		},
		{
			n:        10,
			expected: 5,
		},
	}

	for i, td := range tds {
		r := magicalString(td.n)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
