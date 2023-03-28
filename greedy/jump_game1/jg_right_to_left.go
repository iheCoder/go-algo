package jump_game1

func jump_from_right_left(nums []int) int {
	pos := len(nums) - 1
	steps := 0
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				steps++
				pos = i
				break
			}
		}
	}

	return steps
}
