package clumsy_factorial

func clumsy(n int) int {
	var sub bool
	x := n
	var ans int
	for x > 0 {
		y := x
		if x > 1 {
			y *= x - 1
		}
		if x > 2 {
			y /= x - 2
		}
		if sub {
			ans -= y
		} else {
			ans += y
		}

		sub = true
		x -= 4
	}

	x = n - 3
	for x > 0 {
		ans += x
		x -= 4
	}

	return ans
}
