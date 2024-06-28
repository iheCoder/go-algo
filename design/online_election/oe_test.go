package online_election

import "testing"

func TestOE(t *testing.T) {
	ps := []int{0, 1, 1, 0, 0, 1, 0}
	ts := []int{0, 5, 10, 15, 20, 25, 30}
	c := Constructor(ps, ts)
	//t.Log(c.Q(25))
	ps = []int{0, 0, 0, 0, 1}
	ts = []int{0, 6, 39, 52, 75}
	c = Constructor(ps, ts)
	t.Log(c.Q(78))
}
