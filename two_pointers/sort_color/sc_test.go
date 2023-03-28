package sort_color

import (
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	r    []int
}

func TestChInSlice(t *testing.T) {
	x := []int{1, 2, 3}
	alterSlice(x)
	t.Log(x)
}

func alterSlice(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i]++
	}
}

func TestSCMap(t *testing.T) {
	tds := []testData{
		{
			nums: []int{2, 0, 2, 1, 1, 0},
			r:    []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums: []int{2, 0, 1},
			r:    []int{0, 1, 2},
		},
		{
			nums: []int{1},
			r:    []int{1},
		},
	}

	for _, td := range tds {
		sortColorsMap(td.nums)
		if !reflect.DeepEqual(td.nums, td.r) {
			t.Fatalf("expect %v got %v", td.r, td.nums)
		}
	}
}
