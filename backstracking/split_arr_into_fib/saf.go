package split_arr_into_fib

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(num string) []int {
	n := len(num)
	var ans []int
	var f func(prevSum, last, index int) bool
	f = func(prevSum, last, index int) bool {
		if index == n {
			return len(ans) >= 3
		}

		var cur int
		for i := index; i < n; i++ {
			if i > index && num[index] == '0' {
				break
			}

			cur = cur*10 + int(num[i]-'0')
			if cur > math.MaxInt32 {
				break
			}

			if len(ans) >= 2 {
				if cur > prevSum {
					break
				}
				if cur < prevSum {
					continue
				}
			}

			ans = append(ans, cur)
			if f(last+cur, cur, i+1) {
				return true
			}
			ans = ans[:len(ans)-1]
		}

		return false
	}
	f(0, 0, 0)

	return ans
}

func castToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

// 		if i >= n {
//			return selected
//		}
//
//		if num[i] == '0' {
//			if prevSum == 0 {
//				return f(last, 0, i+1, append(selected, 0))
//			}
//			return nil
//		}
//
//		for j := i + 1; j <= n; j++ {
//			x := castToInt(num[i:j])
//			if x > prevSum {
//				return nil
//			}
//			if x == prevSum {
//				result := f(last+prevSum, prevSum, j, append(selected, x))
//				if result != nil {
//					return result
//				}
//			}
//		}
