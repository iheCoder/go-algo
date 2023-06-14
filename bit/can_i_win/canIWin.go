package can_i_win

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}

	if (maxChoosableInteger*(maxChoosableInteger+1))>>1 < desiredTotal {
		return false
	}

	dp := make([]int8, 1<<maxChoosableInteger)
	// usedNum代表使用的数字
	// curTotal代表使用的数字和
	var dfs func(usedNum, cueTotal int) int8
	dfs = func(usedNum, cueTotal int) int8 {
		if dp[usedNum] != 0 {
			return dp[usedNum]
		}

		for i := 0; i < maxChoosableInteger; i++ {
			// 若当前数字没有被使用
			if usedNum>>i&1 == 0 {
				// 若当前数字和加上选择的当前数字大于期望和或者另一位player失败，那么当前player成功
				if cueTotal+i+1 >= desiredTotal || dfs(usedNum|1<<i, cueTotal+i+1) == 1 {
					dp[usedNum] = 2
					return 2
				}
			}
		}
		dp[usedNum] = 1
		return 1
	}
	return dfs(0, 0) == 2
}

// 列出了player1和player2的所有可能性，但正因为是所有可能性，所以player1必胜
// 	var sum, mask int
//	var dfs func(isPlayer1 bool) bool
//	dfs = func(isPlayer1 bool) bool {
//		for i := 0; i < maxChoosableInteger; i++ {
//			if mask>>i&1 == 0 {
//				sum += i+1
//				mask |= 1 << i
//				if sum >= desiredTotal {
//					return isPlayer1
//				}
//				if dfs(!isPlayer1) {
//					return true
//				}
//				mask &= ^(1 << i)
//				sum -= i+1
//			}
//		}
//		return false
//	}
//	return dfs(true)
