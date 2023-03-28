package intersection_of_two_arrays

import "sort"

func intersection2Map(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	for _, num := range nums1 {
		m1[num] = true
	}

	m2 := make(map[int]bool)
	for _, num := range nums2 {
		m2[num] = true
	}

	r := make([]int, 0)
	for n, _ := range m1 {
		if m2[n] {
			r = append(r, n)
		}
	}

	return r
}

func intersection(nums1 []int, nums2 []int) []int {
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
			for ; i < len(nums1) && nums1[i] == nums1[index]; i++ {
			}
			j++
			for ; j < len(nums2) && nums2[j] == nums1[index]; j++ {
			}
			index++
		}
	}

	return nums1[:index]
}
