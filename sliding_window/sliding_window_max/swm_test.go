package sliding_window_max

import "testing"

func TestSWM(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	t.Log(maxSlidingWindow(nums, k))
}
