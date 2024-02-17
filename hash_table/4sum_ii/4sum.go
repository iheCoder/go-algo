package _sum_ii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1, m2 := make(map[int]int), make(map[int]int)

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m1[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			m2[n3+n4]++
		}
	}

	var ans int
	for k, v := range m1 {
		ans += m2[-k] * v
	}
	return ans
}
