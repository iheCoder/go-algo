package sum_of_square_numbers

import "math"

func judgeSquareSum(c int) bool {
	if c <= 2 {
		return true
	}

	i, j := 0, squareInt(c)
	var sum int
	for i <= j {
		sum = i*i + j*j
		if sum == c {
			return true
		} else if sum > c {
			j--
		} else {
			i++
		}
	}

	return false
}

func squareInt(x int) int {
	return int(math.Sqrt(float64(x)))
}
