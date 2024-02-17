package divide_two_int

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	var rev bool
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	// 取得从divisor到dividend/2区间内的候选者
	cs := []int{divisor}
	for i := divisor; i >= dividend-i; {
		i += i
		cs = append(cs, i)
	}

	var ans int
	for i := len(cs) - 1; i >= 0; i-- {
		if cs[i] >= dividend {
			ans |= 1 << i
			dividend -= cs[i]
		}
	}
	if rev {
		return -ans
	}
	return ans
}

// 取巧的做法还是超出时间限制了
func divideAdd(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var isNeg bool
	if divisor^dividend < 0 {
		isNeg = true
	}

	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}

	var sum, i int
	for sum <= dividend {
		sum += divisor
		i++
	}
	i--
	if isNeg {
		i = -i
	}
	if i < -2147483648 {
		return -2147483648
	} else if i > 2147483647 {
		return 2147483647
	}
	return i
}
