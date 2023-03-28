package rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	if k%n == 0 {
		return
	}
	sourceIndex := 0
	targetIndex := (k + sourceIndex) % n
	source := nums[0]
	var temp int
	var checkLoopIndex int
	var c int
	for c < n {
		temp = nums[targetIndex]
		nums[targetIndex] = source
		sourceIndex = targetIndex
		targetIndex = (k + sourceIndex) % n
		source = temp

		if checkLoopIndex == sourceIndex {
			sourceIndex = checkLoopIndex + 1
			checkLoopIndex = sourceIndex
			source = nums[sourceIndex]
			targetIndex = (k + sourceIndex) % n
		}
		c++
	}
}
