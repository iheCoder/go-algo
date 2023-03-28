package the_lastest_time_to_catch_a_bus

import "testing"

type testDate struct {
	buses      []int
	passengers []int
	cap        int
	expected   int
}

func TestTPRecordAllPossible(t *testing.T) {
	tds := []testDate{
		{
			buses:      []int{10, 20},
			passengers: []int{2, 17, 18, 19},
			cap:        2,
			expected:   16,
		},
		{
			buses:      []int{20, 30, 10},
			passengers: []int{19, 13, 26, 4, 25, 11, 21},
			cap:        2,
			expected:   20,
		},
		{
			buses:      []int{10, 20},
			passengers: []int{1, 2},
			cap:        2,
			expected:   20,
		},
		{
			buses:      []int{1, 2},
			passengers: []int{10, 20},
			cap:        2,
			expected:   2,
		},
		{
			buses:      []int{3},
			passengers: []int{2, 4},
			cap:        2,
			expected:   3,
		},
	}

	for i, td := range tds {
		r := latestTimeCatchTheBusAllpossible(td.buses, td.passengers, td.cap)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestTPRecordCap(t *testing.T) {
	tds := []testDate{
		{
			buses:      []int{10, 20},
			passengers: []int{2, 17, 18, 19},
			cap:        2,
			expected:   16,
		},
		{
			buses:      []int{20, 30, 10},
			passengers: []int{19, 13, 26, 4, 25, 11, 21},
			cap:        2,
			expected:   20,
		},
		{
			buses:      []int{10, 20},
			passengers: []int{1, 2},
			cap:        2,
			expected:   20,
		},
		{
			buses:      []int{1, 2},
			passengers: []int{10, 20},
			cap:        2,
			expected:   2,
		},
		{
			buses:      []int{3},
			passengers: []int{2, 4},
			cap:        2,
			expected:   3,
		},
		{
			buses:      []int{2},
			passengers: []int{2},
			cap:        2,
			expected:   1,
		},
	}

	for i, td := range tds {
		r := latestTimeCatchTheBus(td.buses, td.passengers, td.cap)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
