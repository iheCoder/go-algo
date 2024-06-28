package binary_trees_with_factors

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	dp := make([]int, n)
	res, mod := 0, int(1e9+7)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for left, right := 0, i-1; left <= right; left++ {
			for right >= 0 && arr[left]*arr[right] > arr[i] {
				right--
			}
			for right >= 0 && arr[left]*arr[right] == arr[i] {
				if left == right {
					dp[i] = (dp[i] + dp[left]*dp[right]) % mod
				} else {
					dp[i] = (dp[i] + dp[left]*dp[right]*2) % mod
				}
				right--
			}
		}
		res = (res + dp[i]) % mod
	}

	return res
}
