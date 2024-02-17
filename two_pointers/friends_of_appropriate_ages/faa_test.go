package friends_of_appropriate_ages

import "testing"

type testData struct {
	ages     []int
	expected int
}

func TestFAA(t *testing.T) {
	tds := []testData{
		{
			ages:     []int{16, 16},
			expected: 2,
		},
		{
			ages:     []int{16, 17, 18},
			expected: 2,
		},
		{
			ages:     []int{20, 30, 100, 110, 120},
			expected: 3,
		},
		{
			ages:     []int{7, 18, 19, 23, 14},
			expected: 2,
		},
		{
			ages:     []int{20, 30, 90, 100, 110, 120},
			expected: 6,
		},
		{
			ages:     []int{8, 85, 24, 85, 69},
			expected: 4,
		},
	}

	for i, td := range tds {
		r := numFriendRequests(td.ages)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
