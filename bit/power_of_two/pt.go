package power_of_two

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	return n&(n-1) == 0
}
