package min_moves_to_equal_arr_elements

func minMoves(nums []int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	var ans int
	for _, num := range nums {
		ans += num - minVal
	}
	return ans
}
