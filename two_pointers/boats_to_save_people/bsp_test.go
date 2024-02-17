package boats_to_save_people

import "testing"

type testData struct {
	people   []int
	limit    int
	expected int
}

func TestBSPSort(t *testing.T) {
	tds := []testData{
		{
			people:   []int{1, 2},
			limit:    3,
			expected: 1,
		},
		{
			people:   []int{3, 2, 2, 1},
			limit:    3,
			expected: 3,
		},
		{
			people:   []int{3, 5, 3, 4},
			limit:    5,
			expected: 4,
		},
		{
			people:   []int{3, 5, 4},
			limit:    5,
			expected: 3,
		},
		{
			people:   []int{1, 5, 4},
			limit:    6,
			expected: 2,
		},
	}

	for i, td := range tds {
		r := numRescueBoats(td.people, td.limit)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
