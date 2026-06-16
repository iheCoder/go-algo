package poor_pigs

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	n := minutesToTest/minutesToDie + 1

	pigs, total := 0, 1
	for total < buckets {
		total *= n
		pigs++
	}

	return pigs
}
