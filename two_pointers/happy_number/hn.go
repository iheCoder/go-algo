package happy_number

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	sn := calSqrSum(n)
	fn := calSqrSum(sn)
	for sn != fn {
		if sn == 1 || fn == 1 {
			return true
		}
		sn = calSqrSum(sn)
		fn = calSqrSum(calSqrSum(fn))
	}
	return sn == 1 || fn == 1
}

func calSqrSum(x int) int {
	var sum int
	for x > 0 {
		sum += squaresInt(x % 10)
		x = x / 10
	}
	return sum
}

func squaresInt(x int) int {
	return x * x
}
