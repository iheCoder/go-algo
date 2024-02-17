package arranging_coins

import "math"

func arrangeCoins(n int) int {
	x := int(math.Sqrt(float64(2 * n)))
	if x*(x+1) > 2*n {
		x--
	}
	return x
}
