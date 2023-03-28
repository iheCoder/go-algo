package partition_labels

import (
	"reflect"
	"testing"
)

type testData struct {
	s        string
	expected []int
}

func TestPLMapIndexSlice(t *testing.T) {
	tds := []testData{
		{
			s:        "eccbbbbdec",
			expected: []int{10},
		},
		{
			s:        "ababcbacadefegdehijhklij",
			expected: []int{9, 7, 8},
		},
	}

	for i, td := range tds {
		r := partitionLabels(td.s)
		if !reflect.DeepEqual(r, td.expected) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
