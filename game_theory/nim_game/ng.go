package nim_game

func canWinNim(n int) bool {
	return n%4 != 0
}

// 记忆化搜索+我赢或他输
// 	dp := make(map[int]bool)
//	var f func(rem int) bool
//	f = func(rem int) bool {
//		if rem <= 3 {
//			return true
//		}
//		if win, ok := dp[rem]; ok {
//			return win
//		}
//		for i := 3; i >= 1; i-- {
//			if !f(rem - i) {
//				dp[rem] = true
//				return true
//			}
//		}
//		dp[rem] = false
//		return false
//	}
//	return f(n)

// (n-1)/3&1==0的规律是不对的，这是博弈论
