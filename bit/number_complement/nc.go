package number_complement

func findComplement(num int) int {
	var i int
	for i = 0; i < 32; i++ {
		if num < (1 << i) {
			break
		}
	}

	return ^num & (1<<i - 1)
}
