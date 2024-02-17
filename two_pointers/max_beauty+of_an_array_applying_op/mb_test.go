package max_beauty_of_an_array_applying_op

import "testing"

func TestMB(t *testing.T) {
	nums := []int{4, 6, 1, 2}
	k := 2
	t.Log(maximumBeauty(nums, k))
}
