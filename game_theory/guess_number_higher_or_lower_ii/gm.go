package guess_number_higher_or_lower_ii

func getMoneyAmount(n int) int {
	if n == 1 {
		return 0
	}

	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			f[i][j] = 50000
			if j-i == 1 {
				f[i][j] = i
			}
			// 若k不选择右边界j，那么当j-i==1时，f[i][j]会等于50000，而这是不对的
			// f[n-1][n]当然等于n-1，但若真等于n-1，那么当j为其他值时就又不对了
			for k := i + 1; k < j; k++ {
				cost := k + maxInt(f[i][k-1], f[k+1][j])
				if cost < f[i][j] {
					f[i][j] = cost
				}
			}
		}
	}
	return f[1][n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 	f := make([][]int, n+1)
//	for i := range f {
//		f[i] = make([]int, n+1)
//	}
//
//	for i := n - 1; i >= 1; i-- {
//		for j := i + 1; j <= n; j++ {
//			f[i][j] = j + f[i][j-1]
//			for k := i; k < j; k++ {
//				cost := k + maxInt(f[i][k-1], f[k+1][j])
//				if cost < f[i][j] {
//					f[i][j] = cost
//				}
//			}
//		}
//	}
//	return f[1][n]

// func getMoneyAmount(n int) int {
//	ans := 50000
//	dp := make(map[int]int)
//
//	var f func(start, end, sum int) int
//	f = func(start, end, sum int) int {
//		span := end - start
//		if span <= 1 {
//			return sum
//		}
//		if span == 2 {
//			return sum + start
//		}
//		var maxAmount int
//		if offset, ok := dp[end-start]; ok {
//			return f(start+offset, end, sum+start+offset)
//		}
//
//		mid := start + (end-start)>>1
//		var index int
//		for i := mid; i < end; i++ {
//			ma := maxAmount
//			maxAmount = maxInt(maxAmount, f(1, i, sum+i))
//			maxAmount = maxInt(maxAmount, f(i, end, sum+i))
//			if maxAmount > ma {
//				index = i
//			}
//		}
//		dp[end-start] = index - start
//		return maxAmount
//	}
//	mid := (n + 1) >> 1
//	var maxAmount int
//	for i := mid; i < n; i++ {
//		maxAmount = 0
//		maxAmount = maxInt(maxAmount, f(1, i, i))
//		maxAmount = maxInt(maxAmount, f(i, n, i))
//		ans = minInt(ans, maxAmount)
//	}
//	return ans
//}

// 尝试从中间找，失败！不是所有的最小花费都包含中间值
// 	var f func(start, end, sum int) int
//	f = func(start, end, sum int) int {
//		if end-start <= 1 {
//			return sum
//		}
//		var maxAmount int
//		mid := start + (end-start)>>1
//		maxAmount = maxInt(maxAmount, f(1, mid-1, sum+mid))
//		maxAmount = maxInt(maxAmount, f(mid+1, end, sum+mid))
//		return maxAmount
//	}
//	return f(1, n, 0)

// 最终失败的
//ans := 50000
////dp := make(map[int]int)
//
//// 不能向左右走，而应该向左或者向右走
//var f func(start, end, sum int, isLeft bool) int
//f = func(start, end, sum int, isLeft bool) int {
//	span := end - start
//	if span < 1 {
//		return sum
//	}
//	if span == 1 {
//		return sum + start
//	}
//	var maxAmount int
//	//if offset, ok := dp[span]; ok {
//	//	return f(start+offset, end, sum+start+offset)
//	//}
//
//	mid := start + (end-start)>>1
//	//var index int
//	for i := mid; i < end; i++ {
//		//ma := maxAmount
//		if isLeft {
//			maxAmount = maxInt(maxAmount, f(1, i-1, sum+i, isLeft))
//		} else {
//			maxAmount = maxInt(maxAmount, f(i+1, end, sum+i, isLeft))
//		}
//		//if maxAmount > ma {
//		//	index = i
//		//}
//	}
//	//dp[end-start] = index - start
//	return maxAmount
//}
//mid := (n + 1) >> 1
//var maxAmount int
//for i := mid; i < n; i++ {
//	maxAmount = 0
//	maxAmount = maxInt(maxAmount, f(1, i-1, i, true))
//	maxAmount = maxInt(maxAmount, f(i+1, n, i, false))
//	ans = minInt(ans, maxAmount)
//}
//return ans
