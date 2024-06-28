package peak_index_in_a_mountain

import "testing"

func TestPI(t *testing.T) {
	arr := []int{0, 1, 0}
	t.Log(peakIndexInMountainArray(arr))
}
