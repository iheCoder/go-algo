package find_all_dup_in_an_array

func findDuplicates(nums []int) []int {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	var ans []int
	for i, num := range nums {
		if num != i+1 {
			ans = append(ans, num)
		}
	}
	return ans
}
