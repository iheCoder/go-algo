package gas_stations

import "testing"

func TestGS(t *testing.T) {
	gs := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	t.Log(canCompleteCircuitForce(gs, cost))

	gs = []int{2, 1, 1, 6}
	cost = []int{1, 5, 2, 0}
	t.Log(canCompleteCircuitForce(gs, cost))

	gs = []int{6, 2, 3, 3}
	cost = []int{2, 9, 1, 1}
	t.Log(canCompleteCircuitForce(gs, cost))

}
