package coin_change_ii

import (
	"sort"
)

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	sort.Ints(coins)
	var i int
	for i = len(coins) - 1; i >= 0 && coins[i] > amount; i-- {
	}
	if i < 0 {
		return 0
	}
	coins = coins[:i+1]

	dp := make(map[int]int)
	dp[0] = 1
	for j := len(coins) - 1; j >= 0; j-- {
		v := coins[j]
		dp[v]++
		for k := coins[j] + 1; k <= amount; k++ {
			dp[k] += dp[v] * dp[k-v]
		}
	}

	return dp[amount]
}

//for _, coin := range coins {
//	dp[coin]++
//	for j := coin + 1; j <= amount; j++ {
//		dp[j] += dp[coin] * dp[j-coin]
//	}
//}

//for _, coin := range coins {
//	dp[coin] = 1
//}
//
//for j := 1; j <= amount; j++ {
//	for k := 1; k <= j/2; k++ {
//		dp[j] += dp[k] * dp[j-k]
//	}
//}
