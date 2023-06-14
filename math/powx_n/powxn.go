package powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n&1 == 0 {
		return y * y
	}
	return y * y * x
}

// 超时
// 	if n == 0 {
//		return 1
//	}
//	var isNeg bool
//	if n < 0 {
//		n = -n
//		isNeg = true
//	}
//	y := float64(1)
//	for i := 0; i < n; i++ {
//		y *= x
//	}
//	if isNeg {
//		y = 1 / y
//	}
//	return y
