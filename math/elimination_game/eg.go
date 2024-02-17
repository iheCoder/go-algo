package elimination_game

func lastRemaining(n int) int {
	r := 1
	// k代表轮次，cnt代表总数量，step代表步数
	k, cnt, step := 0, n, 1
	for cnt > 1 {
		if k%2 == 0 {
			r += step
		} else {
			if cnt%2 == 1 {
				r += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return r
}

// 找规律
// 1  1
// 2  2
// 3  2
// 4  2
// 5  2
// 6  4
// 7  4
// 8  6
// 9  6
// 10 8
// 11 8
// 12 6
// 13 6
