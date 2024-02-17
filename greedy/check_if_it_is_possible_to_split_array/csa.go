package check_if_it_is_possible_to_split_array

func canSplitArray(nums []int, m int) bool {
	if len(nums) <= 2 {
		return true
	}

	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i]+nums[i+1] >= m {
			return true
		}
	}
	return false
}
