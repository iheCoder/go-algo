package daily_temperatures

import "testing"

func TestDT(t *testing.T) {
	ts := []int{73, 74, 75, 71, 69, 72, 76, 73}
	t.Log(dailyTemperatures(ts))
}
