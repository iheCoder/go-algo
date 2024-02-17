package nth_digit

import "strconv"

// 这就超出内存限制了啊
// 即使提前终止也还是超出了内存限制
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	var c, i int
	rank := 1
	rankLimit := 10
	for i = 0; i <= n && c < n; {
		i++
		if i == rankLimit {
			rank++
			rankLimit *= 10
		}
		c += rank
	}

	is := strconv.Itoa(i)
	return int(is[len(is)-1-c+n] - '0')
}
