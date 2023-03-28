package _sum

import (
	"reflect"
	"testing"
)

type testData struct {
	nums     []int
	expected [][]int
}

func TestThreeSumMap(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: make([][]int, 0),
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for i, td := range tds {
		r := threeSumIJK(td.nums)
		if !twoDArrayEqual(td.expected, r) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestThreeSum(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: make([][]int, 0),
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for i, td := range tds {
		r := threeSum(td.nums)
		if !twoDArrayEqual(td.expected, r) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestEqual(t *testing.T) {
	a := [][]int{}
	b := make([][]int, 0)
	t.Log(twoDArrayEqual(a, b))

	a = [][]int{{1, 2, 3}}
	b = [][]int{{1, 2, 3}}
	t.Log(twoDArrayEqual(a, b))

	a = [][]int{{1, 2, 3}}
	b = [][]int{{2, 3, 1}}
	t.Log(twoDArrayEqual(a, b))

	a = [][]int{{1, 2, 3}}
	b = [][]int{{2, 3, 2}}
	t.Log(twoDArrayEqual(a, b))

	a = [][]int{{1, 2, 3}}
	b = [][]int{{1, 3, 2}, {1}}
	t.Log(twoDArrayEqual(a, b))

	a = [][]int{{1, 2, 3}}
	b = [][]int{{1, 3, 2}, {}}
	t.Log(twoDArrayEqual(a, b))

	a = [][]int{{1, 2, 3}, {4, 5}}
	b = [][]int{{1, 2, 3}, {5, 4}}
	t.Log(twoDArrayEqual(a, b))
}

func twoDArrayEqual(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	} else if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	c := 0
	for _, ai := range a {
		for _, bi := range b {
			if aInBSlice(ai, bi) {
				c++
				break
			}
		}
	}

	return c == len(a)
}

func aInBSlice(x, y []int) bool {
	if x == nil && y == nil {
		return true
	} else if (x == nil && y != nil) || (x != nil && y == nil) {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	xm := make(map[int]int)
	for _, xi := range x {
		xm[xi]++
	}

	ym := make(map[int]int)
	for _, yi := range y {
		ym[yi]++
	}

	return reflect.DeepEqual(xm, ym)
}
