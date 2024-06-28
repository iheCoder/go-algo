package car_fleet

import "testing"

func TestCF(t *testing.T) {
	pos := []int{6, 8}
	sp := []int{3, 2}
	target := 10
	//t.Log(carFleet(target, pos, sp))
	pos = []int{0, 4, 2}
	sp = []int{2, 1, 3}
	t.Log(carFleet(target, pos, sp))
}
