package jump_game

func canJumpLR(nums []int) bool {
	pos := 0
	for pos < len(nums)-1 {
		if pos+nums[pos] >= len(nums)-1 {
			return true
		}
		var maxStep int
		var maxI int
		// 应该要等于最大step
		for i := 1; i <= nums[pos]; i++ {
			if pos+i >= len(nums)-1 {
				return true
			}
			if nums[pos+i] >= maxStep {
				maxI = i
				maxStep = nums[pos+i]
			}
		}
		if maxStep == 0 {
			return false
		}
		pos += maxI
	}

	return true
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
