package house_robber_iv

import "math"

func minCapability(nums []int, k int) int {
	minX, maxX := math.MaxInt, math.MinInt
	for _, num := range nums {
		if num > maxX {
			maxX = num
		}
		if num < minX {
			minX = num
		}
	}

	isSatisfyK := func(x int) bool {
		var count int
		for i := 0; i < len(nums); i++ {
			if nums[i] <= x {
				count++
				i++
			}
		}

		return count >= k
	}

	l, r, mid := minX, maxX, 0
	for r > l {
		mid = (r-l)/2 + l

		if isSatisfyK(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
