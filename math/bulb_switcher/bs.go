package bulb_switcher

import (
	"math"
	"sort"
)

const maxValue = 31622

func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}

	x := sort.Search(n, func(i int) bool {
		return i*i >= n
	})
	if x*x > n {
		return x - 1
	}
	return x
}

func isPerfectSqrt(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}
