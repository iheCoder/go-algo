package find_1st_and_last_pos_of_ele_in_sorted_arr

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	p := binarySearch(nums, target)
	if p == -1 {
		return []int{-1, -1}
	}
	i, j := p-1, p+1
	for i >= 0 {
		if nums[i] != target {
			break
		}
		i--
	}
	for j < len(nums) {
		if nums[j] != target {
			break
		}
		j++
	}
	return []int{i + 1, j - 1}
}

func binarySearch(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for h > l {
		mid := l + (h-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			h = mid
		}
	}
	if nums[h] != target {
		return -1
	}
	return h
}
