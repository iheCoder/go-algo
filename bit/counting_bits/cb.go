package counting_bits

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	bits := make([]int, n+1)
	border := 0
	for i := 1; i < n+1; i++ {
		if i&(i-1) == 0 {
			border = i
		}
		bits[i] = bits[i-border] + 1
	}
	return bits
}
