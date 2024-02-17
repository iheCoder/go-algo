package walking_robot_simulation

import "testing"

func TestWRS(t *testing.T) {
	cs := []int{4, -1, 4, -2, 4}
	obs := [][]int{{2, 4}}
	//t.Log(robotSim(cs, obs))
	cs = []int{-2, -2, 4}
	obs = [][]int{
		{0, -2},
		{0, 1},
		{0, -5},
	}
	//t.Log(robotSim(cs, obs))
	cs = []int{-2, 4}
	obs = [][]int{
		{-2, 0},
		{1, 0},
		{-3, 0},
	}
	//t.Log(robotSim(cs, obs))
	cs = []int{6, -1, -1, 6}
	obs = [][]int{}
	//t.Log(robotSim(cs, obs))
	// 居然还能0，0 ！！！
	cs = []int{2, 2, 5, -1, -1}
	obs = [][]int{
		{-3, 5},
		{-2, 5},
		{3, 2},
		{5, 0},
		{-2, 0},
		{-1, 5},
		{5, -3},
		{0, 0},
		{-4, 4},
		{-3, 4},
	}
	t.Log(robotSim(cs, obs))
}
