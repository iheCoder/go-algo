package max_sum_of_two_non_overlapping_subarrays

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	var ans int
	for l := 1; l < n; l++ {
		if firstLen <= l && secondLen <= n-l {
			ans = maxInt(ans, maxSum(prefix, 0, l, firstLen)+maxSum(prefix, l, n, secondLen))
		}
		if firstLen <= n-l && secondLen <= l {
			ans = maxInt(ans, maxSum(prefix, 0, l, secondLen)+maxSum(prefix, l, n, firstLen))
		}
	}

	return ans
}

func maxSum(prefix []int, start, end, len int) int {
	var ans int
	for i := start + len; i <= end; i++ {
		ans = maxInt(ans, prefix[i]-prefix[i-len])
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
