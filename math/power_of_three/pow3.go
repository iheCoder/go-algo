package power_of_three

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if n&1 == 0 {
		return false
	}
	for n >= 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return n == 1
}
