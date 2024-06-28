package card_flipping_game

func flipgame(fronts []int, backs []int) int {
	sameSet := make(map[int]bool)
	n := len(fronts)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			sameSet[fronts[i]] = true
		}
	}

	res := 3000
	for i := 0; i < n; i++ {
		if !sameSet[fronts[i]] && res > fronts[i] {
			res = fronts[i]
		}
	}
	for i := 0; i < n; i++ {
		if !sameSet[backs[i]] && res > backs[i] {
			res = backs[i]
		}
	}

	return res % 3000
}
