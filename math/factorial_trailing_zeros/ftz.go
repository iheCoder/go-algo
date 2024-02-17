package factorial_trailing_zeros

func trailingZeroes(n int) int {
	var ans int
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}
