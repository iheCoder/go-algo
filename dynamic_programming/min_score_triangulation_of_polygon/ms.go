package min_score_triangulation_of_polygon

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for d := 2; d < n; d++ {
		for i := 0; i+d < n; i++ {
			j := i + d
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				dp[i][j] = minInt(dp[i][j], dp[i][k]+dp[k][j]+values[i]*values[j]*values[k])
			}
		}
	}

	return dp[0][n-1]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
