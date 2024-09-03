package longest_arithmetic_subsequence

func longestArithSeqLength(nums []int) int {
	minv, maxv := nums[0], nums[0]
	for _, num := range nums {
		if num < minv {
			minv = num
		}
		if num > maxv {
			maxv = num
		}
	}

	diff := maxv - minv
	ans := 1
	for d := -diff; d <= diff; d++ {
		f := make([]int, maxv+1)
		for _, num := range nums {
			prev := num - d
			if prev >= minv && prev <= maxv && f[prev] != 0 {
				f[num] = maxInt(f[num], f[prev]+1)
				ans = maxInt(ans, f[num])
			}
			f[num] = maxInt(f[num], 1)
		}
	}

	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
