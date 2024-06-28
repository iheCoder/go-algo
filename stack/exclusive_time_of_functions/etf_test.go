package exclusive_time_of_functions

import "testing"

func TestETF(t *testing.T) {
	n := 1
	logs := []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"}
	t.Log(exclusiveTime(n, logs))
}
