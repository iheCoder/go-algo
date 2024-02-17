package heaters

import "testing"

type testData struct {
	houses   []int
	heaters  []int
	expected int
}

func TestHeaters(t *testing.T) {
	tds := []testData{
		{
			houses:   []int{1, 5},
			heaters:  []int{2},
			expected: 3,
		},
	}

	for i, td := range tds {
		r := findRadius(td.houses, td.heaters)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
