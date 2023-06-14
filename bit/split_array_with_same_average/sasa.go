package split_array_with_same_average

func splitArraySameAverage(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	n := len(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}
	var isPossible bool
	m := n >> 1
	for i := 1; i < m; i++ {
		if sum*i%n == 0 {
			isPossible = true
			break
		}
	}
	if !isPossible {
		return false
	}

	// dp[i][x]=true代表当前已经遍历过的元素中可以取出i个元素和为x
	dp := make([]map[int]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true

	for _, num := range nums {
		for i := m; i > 0; i-- {
			for j := range dp[i-1] {
				cur := j + num
				if cur*n == sum*i {
					return true
				}
				dp[i][cur] = true
			}
		}
	}
	return false
}

func hasSubarrayWithSum(nums []int, m, sum, start, curSum int) bool {
	if curSum == sum && m == 0 {
		return true
	}
	if curSum > sum || m == 0 || start >= len(nums) {
		return false
	}

	for i := start; i < len(nums); i++ {
		curSum += nums[i]
		if hasSubarrayWithSum(nums, m-1, sum, i+1, curSum) {
			return true
		}
		curSum -= nums[i]
	}
	return false
}

// 判断是否大于sum/2来剪枝是不太对的

// 通过后缀和向后看让大的变小还是比另一个大来达到剪枝的问题在于0到任何正数都是变大的
// 		qa := (sumA + suffixSum[i]) / (float64(n) - lb)
//		qb := sumB / lb
// 通过判断长度不为0解决，然而仍然超出时间限制

// 	var dfs func(sumA, la, sumB, lb int, i int) bool
//	dfs = func(sumA, la, sumB, lb int, i int) bool {
//		if i >= n {
//			if la <= 0 || lb <= 0 {
//				return false
//			}
//			return sumA*lb == sumB*la
//		}
//
//		newSumA := sumA + nums[i]
//		if dfs(newSumA, la+1, sumB, lb, i+1) {
//			return true
//		}
//
//		newSumB := sumB + nums[i]
//		if dfs(sumA, la, newSumB, lb+1, i+1) {
//			return true
//		}
//
//		return false
//	}
//	return dfs(0, 0, 0, 0, 0)

// 96s!!!离谱至极
// 虽然求得了规律sm*n == sum*m，但是通过位来获取所有的组合还是太麻烦的一个过程
// 	var sm, m int
//	for mask := 1; mask < (1<<n)-1; mask++ {
//		sm, m = 0, 0
//		for i := 0; i < n; i++ {
//			if mask>>i&1 > 0 {
//				sm += nums[i]
//				m++
//			}
//		}
//		if sm*n == sum*m {
//			return true
//		}
//	}
//	return false

// 通过回溯法找到所有目标sum的m个数组和也是不行的。用时还是太长了
// 	ms := make([]int, 0)
//	for i := 1; i < n; i++ {
//		if sum*i%n == 0 {
//			ms = append(ms, i)
//		}
//	}
//	if len(ms) == 0 {
//		return false
//	}
//
//	for _, m := range ms {
//		if hasSubarrayWithSum(nums, m, sum*m/n, 0, 0) {
//			return true
//		}
//	}
//
//	return false
