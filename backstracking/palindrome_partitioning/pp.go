package palindrome_partitioning

func partition(s string) [][]string {
	var ans [][]string
	n := len(s)
	var f func(i int, cur []string)
	f = func(i int, cur []string) {
		if i >= n {
			ans = append(ans, cur)
			return
		}

		for j := i + 1; j <= n; j++ {
			if isPalindrome(s[i:j]) {
				curr := make([]string, len(cur))
				copy(curr, cur)
				curr = append(curr, s[i:j])
				f(j, curr)
			}
		}
	}
	f(0, make([]string, 0))
	return ans
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
