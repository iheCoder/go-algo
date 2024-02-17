package find_min_in_rotated_sorted_array

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	l, h := 0, n-1
	var mid int
	// 小中大，那么就是l
	// 大小中，那么就在l、mid之间
	// 中大小，那么就在mid、h之间
	for {
		mid = l + (h-l)/2
		if nums[l] > nums[h] && nums[h] > nums[mid] {
			h = mid
		} else if nums[mid] > nums[l] && nums[l] > nums[h] {
			l = mid
		} else {
			// 除了小中大的情况，还有可能就两个了
			return minInt(nums[l], nums[h])
		}
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
