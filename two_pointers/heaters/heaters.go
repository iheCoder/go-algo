package heaters

import (
	"math"
	"sort"
)

// 完全变成求一个heater的思路了
func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	sort.Ints(houses)

	var i int
	var minDis int
	var r int
	for _, house := range houses {
		for ; i < len(heaters) && heaters[i] < house; i++ {
		}
		minDis = math.MaxInt
		if i > 0 {
			minDis = absInt(house - heaters[i-1])
			if i >= len(heaters) {
				if minDis > r {
					r = minDis
				}
				continue
			}
		}
		minDis = minInt(minDis, absInt(house-heaters[i]))
		if minDis > r {
			r = minDis
		}
	}

	return r
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxIntDiff(x, y int, r int) int {
	d := absoluteInt(x - y)
	return maxInt(d, r)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func absoluteInt(x int) int {
	return int(math.Abs(float64(x)))
}

// 	sort.Ints(heaters)
//	hs, he := heaters[0], heaters[len(heaters)-1]
//	sort.Ints(houses)
//	hss, hse := houses[0], houses[len(houses)-1]
//
//	r := 0
//	r = maxIntDiff(hs, hse, r)
//	r = maxIntDiff(he, hss, r)
//	return r

// heaters 若是只有两个还真可以求两个数的绝对距离
