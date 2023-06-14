package rotate_image

import (
	"strconv"
	"strings"
	"testing"
)

func TestRI(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	t.Log(matrix)
	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(matrix)
	t.Log(matrix)
}

func TestGenMatrixTestCases(t *testing.T) {
	n := 5
	t.Log(wrapToString(genMatrix(n)))
	n = 6
	t.Log(wrapToString(genMatrix(n)))
	n = 7
	t.Log(wrapToString(genMatrix(n)))
	n = 8
	t.Log(wrapToString(genMatrix(n)))
	n = 19
	t.Log(wrapToString(genMatrix(n)))
	n = 20
	t.Log(wrapToString(genMatrix(n)))
}

func wrapToString(r [][]int) string {
	var s string
	for _, ri := range r {
		si := make([]string, 0, len(ri))
		for _, v := range ri {
			si = append(si, strconv.Itoa(v))
		}
		s += "[" + strings.Join(si, ",") + "],"
	}
	s = s[:len(s)-1]
	s = "[" + s + "]"
	return s
}

func genMatrix(n int) [][]int {
	var count int
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		ri := make([]int, n)
		for j := 0; j < n; j++ {
			count++
			ri[j] = count
		}
		r[i] = ri
	}
	return r
}
