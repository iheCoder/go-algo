package the_k_strongest_values_in_an_array

import (
	"math"
	"sort"
)

func getStrongest(arr []int, k int) []int {
	if len(arr) == k {
		return arr
	}

	sort.Ints(arr)
	median := arr[(len(arr)-1)/2]

	i, j := 0, len(arr)-1
	r := make([]int, k)
	var ri int
	for ri < k && i <= j {
		if getLeftMaxAbs(arr[i], arr[j], median) {
			r[ri] = arr[i]
			ri++
			i++
		} else {
			r[ri] = arr[j]
			ri++
			j--
		}
	}

	return r
}

func getLeftMaxAbs(x, y, median int) bool {
	xm := math.Abs(float64(x - median))
	ym := math.Abs(float64(y - median))

	if xm > ym {
		return true
	} else if xm < ym {
		return false
	} else {
		if x > y {
			return true
		}
		return false
	}
}

func findMed(arr []int) int {
	n := len(arr)
	medianIndex := n / 2
	median := arr[medianIndex]
	if n%2 == 0 {
		median = (median + arr[medianIndex-1]) / 2
	}
	return median
}
