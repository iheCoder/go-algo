package jump_game1

func jump(nums []int) int {
	n := len(nums)
	maxPos := 0
	end := 0
	steps := 0
	for i := 0; i < n; i++ {
		maxPos = maxInt(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			steps++
		}
	}

	return steps
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
