package guess_number_higher_or_lower_ii

import "testing"

type testData struct {
	n        int
	expected int
}

func TestGetMoneyAmount(t *testing.T) {
	tds := []testData{
		{
			n:        3,
			expected: 2,
		},
		{
			n:        4,
			expected: 4,
		},
		{
			n:        5,
			expected: 6,
		},
		{
			n:        6,
			expected: 8,
		},
		{
			n:        7,
			expected: 10,
		},
		{
			n:        8,
			expected: 12,
		},
		{
			n:        9,
			expected: 14,
		},
		{
			n:        10,
			expected: 16,
		},
	}
	for i, td := range tds {
		got := getMoneyAmount(td.n)
		if td.expected != got {
			t.Fatalf("%d n %d expected %d got %d", i, td.n, td.expected, got)
		}
	}
}
