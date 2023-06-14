package predict_the_winner

func PredictTheWinner(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	var f func(i, j, score, peerScore int) int
	f = func(i, j, score, peerScore int) int {
		if i > j {
			return score - peerScore
		}

		return iWin(f(i+1, j, peerScore, score+nums[i]), f(i, j-1, peerScore, score+nums[j]))
	}
	return f(0, len(nums)-1, 0, 0) >= 0
}

func iWin(x, y int) int {
	if x < 0 || y < 0 {
		return 1
	}
	if x == 0 || y == 0 {
		return 0
	}
	return -1
}

//if !f(i+1, j, peerScore, score+nums[i]) {
//	return true
//}
//if !f(i, j-1, peerScore, score+nums[j]) {
//	return true
//}
//return false
