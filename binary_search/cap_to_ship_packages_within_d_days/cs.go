package cap_to_ship_packages_within_d_days

import "sort"

func shipWithinDays(weights []int, days int) int {
	var left, right int
	for _, weight := range weights {
		if weight > left {
			left = weight
		}
		right += weight
	}

	return left + sort.Search(right-left, func(i int) bool {
		x := left + i
		var sum int
		d := 1
		for _, weight := range weights {
			if weight+sum > x {
				d++
				sum = 0
			}
			sum += weight
		}
		return d <= days
	})
}
