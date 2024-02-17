package divisor_game

func divisorGame(n int) bool {
	if n == 1 {
		return false
	}

	dp := make(map[int]bool)
	var f func(rem int) bool
	f = func(rem int) bool {
		if rem <= 1 {
			return false
		}

		if win, ok := dp[rem]; ok {
			return win
		}

		ran := rem >> 1
		for i := ran; i > 0; i-- {
			if rem%i != 0 {
				continue
			}
			if !f(rem - i) {
				dp[rem] = true
				return true
			}
		}
		dp[rem] = false
		return false
	}
	return f(n)
}
