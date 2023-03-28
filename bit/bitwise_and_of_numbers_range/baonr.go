package bitwise_and_of_numbers_range

func rangeBitwiseAnd(left int, right int) int {
	l := getBitLevel(right)
	c := right - left

	return 0
}

func getBitLevel(x int) int {
	if x == 0 {
		return 0
	}
	l := 0
	for x != 0 {
		x >>= 1
		l++
	}
	return l
}
