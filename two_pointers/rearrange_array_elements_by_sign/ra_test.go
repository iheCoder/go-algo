package rearrange_array_elements_by_sign

import (
	"reflect"
	"testing"
)

type testData struct {
	nums     []int
	expected []int
}

func TestRA(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{-1, 1},
			expected: []int{1, -1},
		},
	}

	for i, td := range tds {
		r := rearrangeArray(td.nums)
		if !reflect.DeepEqual(td.expected, r) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
