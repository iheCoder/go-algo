package cheapest_flights_within_k_stops

import "testing"

func TestFCP(t *testing.T) {
	n := 4
	fs := [][]int{
		{0, 1, 1},
		{0, 2, 5},
		{1, 2, 1},
		{2, 3, 1},
	}
	src, dst := 0, 3
	k := 1
	t.Log(findCheapestPrice(n, fs, src, dst, k))

}
