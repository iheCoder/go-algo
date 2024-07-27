package broken_calculator

// 2xa+y=b
// 2a=(b-y)/x, min(|x|+|y|)

// 2^x * a>=b, x>=sqrt(b/a)
func brokenCalc(startValue int, target int) int {
	var ans int
	for target > startValue {
		ans++
		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
	}
	return ans + startValue - target
}

// 	if target <= startValue {
//		return startValue - target
//	}
//
//	x := int(math.Log2(float64(target / startValue)))
//	xs := (1 << x) * startValue
//	for xs < target {
//		x++
//		xs = (1 << x) * startValue
//	}
//	diff := xs - target
//	ans := x + diff
//
//	if target%2 == 0 {
//		y := x
//		for target/(1<<y) > startValue {
//			y--
//		}
//		diff = startValue - target/(1<<y)
//		tmp := diff + y
//		if tmp < ans {
//			ans = tmp
//		}
//	}
//
//	return ans
