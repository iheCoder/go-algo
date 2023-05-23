package circular_arr_loop

import "testing"

type testData struct {
	nums     []int
	expected bool
}

func TestCAL(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{2, -1, 1, 2, 2},
			expected: true,
		},
		{
			nums:     []int{-1, -2, -3, -4, -5, 6},
			expected: false,
		},
		{
			nums:     []int{1, -1, 5, 1, 4},
			expected: true,
		},
		{
			nums:     []int{1},
			expected: false,
		},
	}

	for i, td := range tds {
		r := circularArrayLoop(td.nums)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
