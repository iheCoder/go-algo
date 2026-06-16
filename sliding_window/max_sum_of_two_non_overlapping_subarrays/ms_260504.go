package max_sum_of_two_non_overlapping_subarrays

func maxSumTwoNoOverlap260504(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	prefix[0] = nums[0]
	for i := 1; i <= n; i++ {
		prefix[i] += prefix[i-1] + nums[i-1]
	}

	return maxInt(
		helper(prefix, firstLen, secondLen),
		helper(prefix, secondLen, firstLen),
	)
}

func helper(prefix []int, l, r int) int {
	maxL, ans := 0, 0
	n := len(prefix) - 1

	for i := l + r; i <= n; i++ {
		curL := prefix[i-r] - prefix[i-r-l]
		curR := prefix[i] - prefix[i-r]
		maxL = maxInt(maxL, curL)
		ans = maxInt(ans, maxL+curR)
	}

	return ans
}
