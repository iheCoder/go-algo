package k_concatenation_max_sum

const mod = 1e9 + 7

func kConcatenationMaxSum(arr []int, k int) int {
	if k == 1 {
		// 计算普通的最大子数组和
		return maxSubArraySum(arr) % int(mod)
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	twice := append(arr, arr...)
	if sum <= 0 {
		// 如果数组元素总和小于等于0，只需要考虑拼接两次的情况求最大子数组和
		return maxSubArraySum(twice) % int(mod)
	}
	// 先计算拼接两次的最大子数组和
	doubleMax := maxSubArraySum(twice)
	// 再计算一次普通最大子数组和
	singleMax := maxSubArraySum(arr)
	return maxInt(doubleMax+(k-2)*sum, singleMax) % int(mod)
}

func maxSubArraySum(a []int) int {
	maxEndingHere := 0
	maxSoFar := 0
	for _, num := range a {
		maxEndingHere = maxInt(num, maxEndingHere+num)
		maxSoFar = maxInt(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//const MOD = 1000000007
//
//func kConcatenationMaxSum(arr []int, k int) int {
//	n := len(arr)
//	maxSum := maxSumFromArr(arr)
//
//	if k == 1 {
//		return maxSum % MOD
//	}
//
//	var ps, mps int
//	for i := 0; i < n; i++ {
//		ps += arr[i]
//		ps %= MOD
//		mps = maxInt(mps, ps)
//	}
//
//	var ss, mss int
//	for i := n - 1; i >= 0; i-- {
//		ss += arr[i]
//		ss %= MOD
//		mss = maxInt(mss, ss)
//	}
//
//	concatSum := (mps + mss) % MOD
//	maxSum %= MOD
//	if ps >= 0 {
//		x := (((k-2)*ps)%MOD + concatSum) % MOD
//		return maxInt(maxSum, x)
//	}
//
//	return maxInt(maxSum, concatSum)
//}
//
//func maxSumFromArr(arr []int) int {
//	var me, mf int
//	for _, num := range arr {
//		me = maxInt(num, me+num)
//		mf = maxInt(mf, me%MOD)
//	}
//
//	return mf
//}
//
//func maxInt(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
