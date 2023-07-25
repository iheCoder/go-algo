package perfect_squares

var dp = map[int]int{}

func numSquares(n int) int {
	if n == 1 {
		return 1
	}
	dp = make(map[int]int)
	return ns(n)
}

func ns(n int) int {
	if n == 1 {
		return 1
	}
	ans := 1<<31 - 1
	var square int
	x := n >> 1
	var y int
	var ok bool
	for x >= 1 {
		square = x * x
		x--
		if square > n {
			continue
		}
		if square == n {
			return 1
		}
		if y, ok = dp[n-square]; !ok {
			y = ns(n - square)
			dp[n-square] = y
		}
		ans = minInt(ans, y+1)
	}
	return ans
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
