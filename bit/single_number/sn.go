package single_number

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}
