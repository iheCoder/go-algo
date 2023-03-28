package _sum_closet

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var l, r int
	var x int
	R := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l = i + 1
		r = len(nums) - 1
		for l < r {
			x = nums[i] + nums[l] + nums[r]
			R = minInt(R, x, target)
			if x == target {
				return target
			} else if x < target {
				l++
			} else {
				r--
			}
		}
	}

	return R
}

func minInt(x, y, target int) int {
	xt := math.Abs(float64(x - target))
	yt := math.Abs(float64(y - target))
	if xt > yt {
		return y
	} else {
		return x
	}
}
