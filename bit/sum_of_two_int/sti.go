package sum_of_two_int

func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}

//func getSum(a int, b int) int {
//	switch {
//	case a > 0 && b > 0:
//		return (a ^ b) | ((a & b) << 1)
//	case a < 0 && b < 0:
//		a, b = -a, -b
//		return -((a ^ b) | ((a & b) << 1))
//	case a > 0 && b < 0:
//		b = -b
//		var r int
//		if b > a {
//			r = -((a ^ b) & (b | a>>1))
//		} else if b == a {
//			return 0
//		} else {
//			r = (a ^ b) & (a | b>>1)
//		}
//		return r
//	default:
//		a = -a
//		var r int
//		if a > b {
//			r = -((a ^ b) & (a | b>>1))
//		} else if b == a {
//			return 0
//		} else {
//			r = (a ^ b) & (b | a>>1)
//		}
//		return r
//	}
//}
