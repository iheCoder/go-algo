package task_scheduler

import "testing"

func TestTS(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2
	t.Log(leastIntervalFetchMax(tasks, n))
	//t.Log(447+429.77+44+298+198+49+12.64+14.3+14.65+15.25)
}
