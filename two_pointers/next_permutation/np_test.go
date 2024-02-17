package next_permutation

import (
	"reflect"
	"sort"
	"testing"
)

type testData struct {
	nums     []int
	expected []int
}

func TestNP(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{1, 5, 1},
			expected: []int{5, 1, 1},
		},
		{
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			nums:     []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			nums:     []int{2, 1},
			expected: []int{1, 2},
		},
		{
			nums:     []int{2, 1, 3, 3},
			expected: []int{2, 3, 1, 3},
		},
	}

	for i, td := range tds {
		nextPermutation(td.nums)
		if !reflect.DeepEqual(td.expected, td.nums) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, td.nums)
		}
	}
}

func TestSortIntsSearch(t *testing.T) {
	arr := []int{1, 2, 3}
	r := sort.SearchInts(arr, 4)
	t.Log(r)
}
