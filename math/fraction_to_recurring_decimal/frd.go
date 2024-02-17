package fraction_to_recurring_decimal

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	var isNeg bool
	if numerator < 0 {
		numerator = -numerator
		isNeg = !isNeg
	}
	if denominator < 0 {
		denominator = -denominator
		isNeg = !isNeg
	}

	rem := numerator
	nums := make([]byte, 0)
	if isNeg {
		nums = append(nums, '-')
	}
	nums = append(nums, strconv.Itoa(rem/denominator)...)
	rem = rem % denominator
	nums = append(nums, '.')

	m := make(map[int]int)
	for rem > 0 && m[rem] == 0 {
		m[rem] = len(nums)
		rem *= 10
		nums = append(nums, '0'+byte(rem/denominator))
		rem = rem % denominator
	}

	sp := m[rem]
	if rem > 0 {
		nums = append(nums[:sp], append([]byte{'('}, nums[sp:]...)...)
		nums = append(nums, ')')
	}

	return string(nums)
}

// 	for rem > 0 {
//		c := 0
//		rem *= 10
//		for rem < denominator {
//			rem *= 10
//			c++
//		}
//		if m[rem] > 0 {
//			sp = m[rem]
//			break
//		} else if m[rem] == 0 {
//			m[rem] = len(nums)
//		}
//		for c > 0 {
//			nums = append(nums, 0)
//			c--
//		}
//		nums = append(nums, rem/denominator)
//		rem = rem % denominator
//	}
