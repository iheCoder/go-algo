package max_distance_to_closest_person

func maxDistToClosest(seats []int) int {
	n := len(seats)
	a := make([]int, 0, n)
	for i, seat := range seats {
		if seat == 1 {
			a = append(a, i)
		}
	}

	var ans, last int
	for i, b := range a {
		if i == 0 && b != 0 {
			ans = b
		} else {
			ans = maxInt(ans, (b-last)/2)
		}
		last = b
	}
	if a[len(a)-1] != n-1 {
		ans = maxInt(ans, n-1-last)
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
