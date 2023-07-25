package coin_change

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	maxValue := 1<<32 - 1
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = maxValue
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i < coins[j] {
				continue
			}
			dp[i] = minInt(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] >= maxValue {
		return -1
	}
	return dp[amount]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//
//	sort.Ints(coins)
//	var i int
//	for i = len(coins) - 1; i >= 0 && coins[i] > amount; i-- {
//	}
//	coins = coins[:i+1]
//
//	if len(coins) == 0 {
//		return -1
//	}
//
//	maxValue := 1<<32 - 1
//	m := make(map[int]int)
//	var f func(i, n int) int
//	f = func(i, n int) int {
//		if i < 0 {
//			m[n] = maxValue
//			return maxValue
//		}
//		var k int
//		v := maxValue
//		for j := i; j >= 0; j-- {
//			x := n
//			c := 0
//			c += x / coins[j]
//			// 不一定只能完全整除啊！比如3、4，amount=10
//			x = x % coins[j]
//			if x == 0 {
//				v = minInt(v, c)
//				continue
//			}
//			if y, ok := m[x]; ok {
//				v = minInt(v, c+y)
//				continue
//			}
//			for k = j - 1; k >= 0 && coins[k] > x; k-- {
//			}
//			for k >= 0 {
//				v = minInt(v, c+f(k, x))
//				k--
//			}
//		}
//		m[n] = v
//		return v
//	}
//	ans := f(len(coins)-1, amount)
//	if ans >= maxValue {
//		return -1
//	}
//	return ans
//}
