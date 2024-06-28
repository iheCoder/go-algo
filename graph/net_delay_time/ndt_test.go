package net_delay_time

import "testing"

func TestNdt(t *testing.T) {
	times := [][]int{
		{1, 2, 1},
	}
	n := 2
	k := 1
	t.Log(networkDelayTime(times, n, k))
}
