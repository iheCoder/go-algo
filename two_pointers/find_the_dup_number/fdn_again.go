package find_the_dup_number

func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = nums[0]
	fast = nums[fast]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}
