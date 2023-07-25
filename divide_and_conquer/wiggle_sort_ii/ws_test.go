package wiggle_sort_ii

import "testing"

func TestWS(t *testing.T) {
	nums := []int{1, 2, 3}
	//wiggleSort(nums)
	//t.Log(nums)
	//nums = []int{1, 2, 3, 4}
	//wiggleSort(nums)
	//t.Log(nums)
	//nums = []int{1, 2, 3, 4, 5}
	//wiggleSort(nums)
	//t.Log(nums)
	//nums = []int{1, 2, 3, 4, 5, 6}
	//wiggleSort(nums)
	//t.Log(nums)
	nums = []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums)
	t.Log(nums)
}
