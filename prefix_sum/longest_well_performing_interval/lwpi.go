package longest_well_performing_interval

func longestWPI(hours []int) int {
	f := func(h int) int {
		if h > 8 {
			return 1
		}

		return -1
	}

	var maxInterval, sum int
	m := make(map[int]int)
	for i, hour := range hours {
		sum += f(hour)
		if sum > 0 {
			maxInterval = maxInt(maxInterval, i+1)
		} else {
			j, ok := m[sum-1]
			if ok {
				maxInterval = maxInt(maxInterval, i-j)
			}
		}

		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
	}

	return maxInterval
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// 非滑动窗口
// func longestWPI(hours []int) int {
//	var maxInterval int
//	var l, r int
//	var curSum int
//
//	f := func(h int) int {
//		if h > 8 {
//			return 1
//		}
//		return -1
//	}
//
//	for r < len(hours) {
//		curSum += f(hours[r])
//		if curSum > 0 {
//			maxInterval = maxInt(maxInterval, r-l+1)
//			r++
//			continue
//		}
//
//		for curSum <= 0 && l <= r {
//			curSum -= hours[l]
//			l++
//		}
//
//		r++
//	}
//
//	return maxInterval
//}
//
//func maxInt(x, y int) int {
//	if x > y {
//		return x
//	}
//
//	return y
//}
