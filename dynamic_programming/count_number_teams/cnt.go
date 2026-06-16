package count_number_teams

func numTeams(rating []int) int {
	var res int
	n := len(rating)
	f := func(j int, aes bool) int {
		var px, py int
		if aes {
			for x := 0; x < j; x++ {
				if rating[x] < rating[j] {
					px++
				}
			}
			for y := j + 1; y < n; y++ {
				if rating[y] > rating[j] {
					py++
				}
			}

		} else {
			for x := 0; x < j; x++ {
				if rating[x] > rating[j] {
					px++
				}
			}
			for y := j + 1; y < n; y++ {
				if rating[y] < rating[j] {
					py++
				}
			}
		}

		return px * py
	}

	for i := 1; i < n-1; i++ {
		res += f(i, true)
		res += f(i, false)
	}

	return res
}
