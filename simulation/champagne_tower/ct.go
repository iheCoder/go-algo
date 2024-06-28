package champagne_tower

import "math"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	row := []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		nextRow := make([]float64, len(row)+1)
		for j, v := range row {
			if v > 1 {
				nextRow[j] += (v - 1) / 2
				nextRow[j+1] += (v - 1) / 2
			}
		}
		row = nextRow
	}
	return math.Min(row[query_glass], 1.)
}

// 还是复杂度太高，超时
// func champagneTower(poured int, query_row int, query_glass int) float64 {
//	total := getCombinationCount(query_row)
//	if total >= poured {
//		return 0
//	}
//
//	//n := query_row + (query_row + 1)
//
//	var f func(x, y int) float64
//	f = func(x, y int) float64 {
//		if y < query_row-x || y > query_row+x {
//			return 0
//		}
//		if x == 0 {
//			return float64(poured) - 1
//		}
//
//		left := f(x-1, y-1)
//		//fmt.Printf("%d %d %f\n", x-1, y-1, left)
//		right := f(x-1, y+1)
//		//fmt.Printf("%d %d %f\n", x-1, y+1, right)
//		return (left+right)/2 - 1
//	}
//
//	return minInt(1+f(query_row, query_glass*2), 1.)
//}
//
//func getCombinationCount(n int) int {
//	return (n * (n + 1)) / 2
//}
//
//func minInt(x, y float64) float64 {
//	if x < y {
//		return x
//	}
//	return y
//}

// 显然dfs复杂度太高了
// 	total := getCombinationCount(query_row)
//	if total > poured {
//		return 0
//	}
//
//	n := query_row + (query_row + 1)
//	dp := make([][]float64, query_row+1)
//	for i := 0; i <= query_row; i++ {
//		dp[i] = make([]float64, n)
//	}
//
//	var dfs func(x, y int, rem float64)
//	dfs = func(x, y int, rem float64) {
//		if rem <= 0 || x > query_row {
//			return
//		}
//
//		dp[x][y] += rem
//		if dp[x][y] > 1 {
//			rem = dp[x][y] - 1
//			dfs(x+1, y-1, rem/2)
//			dfs(x+1, y+1, rem/2)
//		}
//	}
//	dfs(0, n/2, float64(poured))
//
//	return dp[query_row][query_glass*2]

// total := getCombinationCount(query_row)
//if total > poured {
//    return 0
//}
//
//total = total + query_glass + 1 + 1
//dp := make([]float64, total)
//var dfs func(x int, rem float64)
//dfs = func(x int, rem float64) {
//    if rem <= 0 || x >= total {
//       return
//    }
//
//    dp[x] += rem
//    if dp[x] > 1 {
//       rem = dp[x] - 1
//       dfs(2*x, rem/2)
//       dfs(2*x+1, rem/2)
//    }
//}
//dfs(1, float64(poured))
//
//return dp[2*query_row+query_glass]
