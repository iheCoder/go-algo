package rotate_array

import (
	"reflect"
	"testing"
)

type testData struct {
	arr      []int
	k        int
	expected []int
}

func TestRATP(t *testing.T) {
	tds := []testData{
		{
			arr:      []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
	}

	for i, td := range tds {
		rotate(td.arr, td.k)
		if !reflect.DeepEqual(td.expected, td.arr) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, td.arr)
		}
	}
}
