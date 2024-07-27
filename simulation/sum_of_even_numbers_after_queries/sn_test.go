package sum_of_even_numbers_after_queries

import "testing"

func TestSN(t *testing.T) {
	qs := [][]int{
		{1, 0},
		{-3, 1},
		{-4, 0},
		{2, 3},
		{1, 1},
		{2, 2},
	}
	nums := []int{1, 2, 3, 4}
	t.Log(sumEvenAfterQueries(nums, qs))
}

func TestNegMode(t *testing.T) {
	t.Log(-1 % 2)
	t.Log(-3 % 2)
	t.Log(-2 % 2)
}
