package stone_game

import (
	"fmt"
)

func stoneGame(piles []int) bool {
	n := len(piles)

	m := make(map[string]int)
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == j {
			return piles[i]
		}
		k := getKey(i, j)
		if x, ok := m[k]; ok {
			return x
		}

		m[k] = maxInt(piles[i]-f(i+1, j), piles[j]-f(i, j-1))
		return m[k]
	}
	return f(0, n-1) > 0
}

func getKey(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, n)
//		dp[i][i] = piles[i]
//	}
//	//
//	//for i := 0; i < n; i++ {
//	//	for j := i + 1; j < n; j++ {
//	//		dp[i][j] = maxInt(piles[i]+dp[i-1][j], piles[j]+dp[i][j+1])
//	//	}
//	//}
// 	var il, jr int
//	for i := 0; i < n; i++ {
//		for j := n - 1; j >= i; j-- {
//			il, jr = 0, 0
//			if i > 0 {
//				il = dp[i-1][j]
//			}
//			if j < n-1 {
//				jr = dp[i][j+1]
//			}
//			if (j-i)&1 > 0 {
//				dp[i][j] = maxInt(piles[i]+il, piles[j]+jr)
//			} else {
//				dp[i][j] = minInt(il-piles[i], jr-piles[j])
//			}
//		}
//		if dp[i][i] > 0 {
//			return true
//		}
//	}
//
//	return false
