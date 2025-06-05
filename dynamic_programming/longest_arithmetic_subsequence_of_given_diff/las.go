package longest_arithmetic_subsequence_of_given_diff

func longestSubsequence(arr []int, difference int) int {
	m := make(map[int]int)
	m[arr[0]] = 1

	n := len(arr)
	ans := 1
	for i := 1; i < n; i++ {
		if m[arr[i]-difference] > 0 {
			m[arr[i]] = maxInt(m[arr[i]], m[arr[i]-difference]+1)
		}

		m[arr[i]] = maxInt(m[arr[i]], 1)
		ans = maxInt(m[arr[i]], ans)
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
