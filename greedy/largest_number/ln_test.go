package largest_number

import "testing"

func TestLN(t *testing.T) {
	nums := []int{9, 901, 8, 100, 0}
	t.Log(largestNumber(nums))
}
