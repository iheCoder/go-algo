package max_length_of_repeated_subarray

func findLength(nums1 []int, nums2 []int) int {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	n := len(nums2)

	// 1. 从n到1找到重复的字符串，如找到一个，那么就直接返回
	// 在此过程中统计相同的最大值
	var r int
	for i := 0; i < n; i++ {
		l := n - i
		if l == r {
			return r
		}
		for j := 0; j < len(nums1); j++ {
			for steps := 0; steps < l && j+steps < len(nums1); steps++ {
				if nums1[j+steps] != nums2[i+steps] {
					break
				}
				r = maxInt(r, steps+1)
			}
			if l == r {
				return r
			}
		}
	}

	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
