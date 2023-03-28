package happy_number

import "testing"

type testData struct {
	n        int
	expected bool
}

func TestHNFastSlow(t *testing.T) {
	tds := []testData{
		//{
		//	n:        10,
		//	expected: true,
		//},
		//{
		//	n:        100000,
		//	expected: true,
		//},
		{
			n:        8372980,
			expected: false,
		},
	}

	for i, td := range tds {
		r := isHappy(td.n)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
