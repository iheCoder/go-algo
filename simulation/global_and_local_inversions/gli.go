package global_and_local_inversions

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	rightMin := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > rightMin {
			return false
		}
		if nums[i] < rightMin {
			rightMin = nums[i]
		}
	}
	return true
}
