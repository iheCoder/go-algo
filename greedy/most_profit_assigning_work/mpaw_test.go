package most_profit_assigning_work

import "testing"

func TestMpaw(t *testing.T) {
	ds := []int{2, 4, 6, 8, 10}
	ps := []int{10, 20, 30, 40, 50}
	ws := []int{4, 5, 6, 7}
	//t.Log(maxProfitAssignment(ds, ps, ws))
	ds = []int{23, 30, 35, 35, 43, 46, 47, 81, 83, 98}
	ps = []int{8, 11, 11, 20, 33, 37, 60, 72, 87, 95}
	// 87+37+60+87+0+20(11)+95+60+20+87
	ws = []int{95, 46, 47, 97, 11, 35, 99, 56, 41, 92}
	t.Log(maxProfitAssignment(ds, ps, ws))
}
