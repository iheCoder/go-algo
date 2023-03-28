package jump_game

func canJump(nums []int) bool {
	rightMost := 0
	for i, num := range nums {
		if i > rightMost {
			return false
		}

		rightMost = maxInt(rightMost, i+num)
		if rightMost > len(nums)-1 {
			return true
		}
	}

	return false
}
