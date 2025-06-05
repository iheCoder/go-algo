package closest_divisors

import "math"

func closestDivisors(num int) []int {
	x1, y1 := divide(num + 1)
	x2, y2 := divide(num + 2)

	if abs(x1, y1) < abs(x2, y2) {
		return []int{x1, y1}
	}

	return []int{x2, y2}
}

func abs(x, y int) int {
	z := x - y
	if z < 0 {
		return -z
	}

	return z
}

func divide(n int) (int, int) {
	for x := int(math.Sqrt(float64(n))); x >= 0; x-- {
		if n%x == 0 {
			return x, n / x
		}
	}

	return 1, n
}
