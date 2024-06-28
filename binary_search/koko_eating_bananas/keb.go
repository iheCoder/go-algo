package koko_eating_bananas

import "sort"

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	mx := piles[len(piles)-1]
	return 1 + sort.Search(mx-1, func(speed int) bool {
		speed++
		var t int
		for _, pile := range piles {
			t += (pile + speed - 1) / speed
		}

		return t <= h
	})
}
