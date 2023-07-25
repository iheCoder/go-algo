package count_num_with_unique_digits

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ans, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}
	return ans
}

func sc9(x int) int {
	y := 1
	for i := 0; i < x; i++ {
		y *= 9 - i
	}
	return y
}

func pow(m, n int) int {
	//if n <= 0 {
	//	return 0
	//}
	y := 1
	for i := 0; i < n; i++ {
		y *= m
	}
	return y
}

func combination(m, n int) int {
	if m == n {
		return 1
	}
	if m < n {
		return 0
	}
	return fac(m) / (fac(n) * fac(m-n))
}

var facCache = make(map[int]int)

func fac(x int) int {
	if y, ok := facCache[x]; ok {
		return y
	}

	y := 1
	for i := 1; i <= x; i++ {
		y *= i
	}
	facCache[x] = y
	return y
}

// 组合方法求重复的宣告失败
// 	ans := pow(10, n)
//	// 如何处理0不unique的情况呢？
//	// 在n中取两个为相同的数，其他位置可以为任何除0之外的数
//	// + combination(n, j)*pow(9, n-j) - sc9(n-j)
//	for j := 2; j <= n; j++ {
//		ans -= combination(n, j)*sc9(n-j)*10 - (n-j+1)*sc9(n-j)
//	}
//	return ans

// 如何处理1122的情况呢？11可以生成，22也可以生成
// 	ans -= combination(n, 2)*pow(9, n-2)*10 - combination(n-1, 2)*pow(10, n-2)
