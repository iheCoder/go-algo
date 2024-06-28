package reach_a_number

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	var s, n int
	for s < target || (s-target)%2 == 1 {
		n++
		s += n
	}
	return n
}
