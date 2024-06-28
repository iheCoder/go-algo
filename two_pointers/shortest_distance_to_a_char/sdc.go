package shortest_distance_to_a_char

import "math"

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = math.MaxInt
	}

	for i := 0; i < n; i++ {
		j := i
		for ; j < n && s[j] != c; j++ {
		}
		if j >= n {
			break
		}
		ans[i] = abs(i - j)
	}

	for i := n - 1; i >= 0; i-- {
		j := i
		for ; j >= 0 && s[j] != c; j-- {
		}
		if j < 0 {
			break
		}
		ans[i] = minInt(ans[i], abs(i-j))
	}
	return ans
}

func abs(x int) int {
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
