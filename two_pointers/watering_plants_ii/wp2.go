package watering_plants_ii

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	if len(plants) <= 1 {
		return 0
	}

	i, j := 0, len(plants)-1
	wa, wb := capacityA, capacityB
	var r int
	for i < j {
		if wa < plants[i] {
			wa = capacityA
			r++
		}
		wa -= plants[i]
		if wb < plants[j] {
			wb = capacityB
			r++
		}
		wb -= plants[j]
		i++
		j--
	}
	if i == j {
		if wa < plants[i] && wb < plants[j] {
			r++
		}
	}

	return r
}
