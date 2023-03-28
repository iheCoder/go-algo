package string_compress

import (
	"reflect"
	"testing"
)

type testData struct {
	chars    []byte
	expected []byte
}

func TestCompressS(t *testing.T) {
	tds := []testData{
		//{
		//	chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		//	expected: []byte{'a', '2', 'b', '2', 'c', '3'},
		//},
		//{
		//	chars:    []byte{'a'},
		//	expected: []byte{'a'},
		//},
		//{
		//	chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		//	expected: []byte{'a', 'b', '1', '2'},
		//},
		{
			chars:    []byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c'},
			expected: []byte{'a', 'b', 'c', '6'},
		},
		//{
		//	chars:    []byte{'a', 'b', 'c'},
		//	expected: []byte{'a', 'b', 'c'},
		//},
	}

	for i, td := range tds {
		compress(td.chars)
		if !reflect.DeepEqual(td.expected, td.chars) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, td.chars)
		}
	}
}

func TestChangeValInSlice(t *testing.T) {
	cs := []int{1}
	chForIndex(cs)
	t.Log(cs)
	replaceWithAnotherSlice(cs)
	t.Log(cs)

	cs = []int{1, 2, 3}
	appendDecreaseCs(cs)
	t.Log(cs)
}

func chForIndex(cs []int) {
	cs[0] = 100
}

func appendDecreaseCs(cs []int) {
	copy(cs, []int{9, 8, 7})
}

func replaceWithAnotherSlice(cs []int) {
	cs = append(cs[:0], 101, 102, 103, 104, 105)
}
