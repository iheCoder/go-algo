package intersection_of_two_arrays_ii

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i, j int
	var index int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[index] = nums1[i]
			i++
			j++
			index++
		}
	}

	return nums1[:index]
}
