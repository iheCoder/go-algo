package removeElement

import "testing"

func TestRmEle(t *testing.T) {
	type test struct {
		nums []int
		val  int
		r    int
	}
	ts := []test{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			r:    2,
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			r:    5,
		},
		{
			nums: []int{1},
			val:  1,
			r:    0,
		},
		{
			nums: []int{1},
			val:  2,
			r:    1,
		},
		{
			nums: []int{},
			val:  1,
			r:    0,
		},
	}

	for _, ti := range ts {
		r := removeElement(ti.nums, ti.val)
		if ti.r != r {
			t.Fatalf("rm ele expect %d but got %d for arr %v", ti.r, r, ti.nums)
		}
		if !checkNumExpect(ti.nums, r, ti.val) {
			t.Fatalf("rm ele but got %v", ti.nums)
		}
	}
}

func checkNumExpect(nums []int, l int, val int) bool {
	for i := 0; i < l; i++ {
		if nums[i] == val {
			return false
		}
	}
	return true
}
