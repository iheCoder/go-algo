package largest_number

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		xs, ys := 10, 10
		x, y := nums[i], nums[j]
		for xs <= x {
			xs *= 10
		}
		for ys <= y {
			ys *= 10
		}
		return x*ys+y > y*xs+x
	})

	if nums[0] == 0 {
		return "0"
	}

	var r string
	for i := 0; i < len(nums); i++ {
		r += strconv.Itoa(nums[i])
	}
	return r
}

// 本来想通过x求所在位数最大值与x的差来判断，但是发现并不对。比如100和0，0到9只差9，但是100到999却差899，但显然100更应该在前
