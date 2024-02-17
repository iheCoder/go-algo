package subarray_product_less_than_k

import "testing"

func TestSpltk(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	k := 100
	//t.Log(numSubarrayProductLessThanKPrefixProduct(nums, k))
	nums = []int{1, 1, 1}
	k = 2
	//t.Log(numSubarrayProductLessThanKPrefixProduct(nums, k))
	//t.Log(numSubarrayProductLessThanK(nums, k))
	nums = []int{100, 2, 3, 4, 100, 5, 6, 7, 100}
	k = 100
	t.Log(numSubarrayProductLessThanK(nums, k))
}
