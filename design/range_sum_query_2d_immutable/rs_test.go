package range_sum_query_2d_immutable

import "testing"

func TestRS(t *testing.T) {
	m := [][]int{
		{
			3, 0, 1, 4, 2,
		},
		{
			5, 6, 3, 2, 1,
		},
		{
			1, 2, 0, 1, 5,
		},
		{
			4, 1, 0, 1, 7,
		},
		{
			1, 0, 3, 0, 5,
		},
	}
	nm := Constructor(m)
	t.Log(nm.SumRegion(1, 1, 2, 2))
}
