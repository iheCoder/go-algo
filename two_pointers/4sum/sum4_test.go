package _sum

import "testing"

type testData struct {
	nums     []int
	target   int
	expected [][]int
}

func TestSum4(t *testing.T) {
	tds := []testData{
		{
			nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
			target: -11,
		},
		{
			nums:   []int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5},
			target: 0,
		},
		//{
		//	nums:     []int{1, 4, 5, 2, 0},
		//	target:   7,
		//	expected: [][]int{{0, 1, 2, 4}},
		//},
		//{
		//	nums:     []int{2, 2, 2, 2, 2},
		//	target:   8,
		//	expected: [][]int{{2, 2, 2, 2}},
		//},
		//{
		//	nums:     []int{1, 0, -1, 0, -2, 2},
		//	target:   0,
		//	expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		//},
	}

	for _, td := range tds {
		r := fourSum(td.nums, td.target)
		t.Log(r)
		//if r != td.expected {
		//	t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		//}
	}
}
