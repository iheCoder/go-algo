package hamming_distance

func hammingDistance(x int, y int) int {
	if x == y {
		return 0
	}

	r := x ^ y
	var count int
	for i := 0; i < 32; i++ {
		if r>>i&1 > 0 {
			count++
		}
	}
	return count
}
