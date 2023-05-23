package successful_pairs_of_spells_and_potions

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	r := make([]int, len(spells))

	for i := 0; i < len(spells); i++ {
		r[i] = len(potions) - sort.SearchInts(potions, floorInt(success, spells[i]))
	}

	return r
}

func floorInt(x int64, y int) int {
	z := float64(x) / float64(y)
	intZ := int(z)
	if z > float64(intZ) {
		return intZ + 1
	}
	return intZ
}
