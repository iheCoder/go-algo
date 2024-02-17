package find_k_closest_elements

import "testing"

func TestFCE(t *testing.T) {
	// 找到的i可不一定是需要的喔
	arr := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	r := findClosestElements(arr, 3, 5)
	t.Log(r)
}
