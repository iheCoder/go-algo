package advantage_shuffle

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	if len(nums1) <= 1 {
		return nums1
	}

	r := make([]int, 0, len(nums2))
	sort.Ints(nums1)
	var index int
	for _, num := range nums2 {
		index = sort.SearchInts(nums1, num)
		if index < len(nums1) && nums1[index] == num {
			index++
			for ; index < len(nums1) && nums1[index] == num; index++ {
			}
		}
		if index >= len(nums1) {
			index = 0
		}
		r = append(r, nums1[index])
		nums1 = append(nums1[:index], nums1[index+1:]...)
	}

	return r
}
