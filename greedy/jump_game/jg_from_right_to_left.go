package jump_game

func canJumpRL(nums []int) bool {
	n := len(nums) - 1
	i := n
	var j int
	found := false
	for i > 0 {
		for j = 0; j < i; j++ {
			if nums[j] >= i-j {
				i = j
				found = true
				break
			}
		}
		if !found {
			return false
		}
		found = false
	}
	return true
}
