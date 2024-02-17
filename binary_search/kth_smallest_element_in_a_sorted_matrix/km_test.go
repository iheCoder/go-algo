package kth_smallest_element_in_a_sorted_matrix

import "testing"

func TestKM(t *testing.T) {
	m := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 8
	t.Log(kthSmallest(m, k))
}
