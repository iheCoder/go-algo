package knight_dialer

var jumps = [][]int{
	{4, 6},
	{6, 8},
	{7, 9},
	{4, 8},
	{0, 3, 9},
	{},
	{0, 1, 7},
	{2, 6},
	{1, 3},
	{2, 4},
}

func knightDialer(n int) int {
	const mod int = 1_000_000_007

	dp := make([]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = 1
	}

	c := 1
	for c < n {
		p := make([]int, 10)
		for i, count := range dp {
			for _, j := range jumps[i] {
				p[j] += count
				p[j] %= mod
				// total放里面累加自然会更大啊！
			}
		}

		c++
		dp = p
	}

	var total int
	for _, num := range dp {
		total = (total + num) % mod
	}

	return total % mod
}
