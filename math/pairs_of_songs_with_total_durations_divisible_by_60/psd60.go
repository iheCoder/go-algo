package pairs_of_songs_with_total_durations_divisible_by_60

func numPairsDivisibleBy60(time []int) int {
	m := make([]int, 60)
	for _, x := range time {
		m[x%60]++
	}

	var ans int
	for i := 0; i < 60; i++ {
		if m[i] <= 0 {
			continue
		}

		x := (60 - i) % 60
		if m[x] <= 0 {
			continue
		}

		if i == x {
			ans += (m[i] * (m[i] - 1)) / 2
		} else {
			ans += m[i] * m[x]
		}

		m[i] = 0
		m[x] = 0
	}
	return ans
}

func cnm(n, m int) int {
	return m * n
}
