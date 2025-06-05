package break_a_palindrome

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}

	n := len(palindrome)
	var flag bool
	ans := []byte(palindrome)
	for i := 0; i < n; i++ {
		if ans[i] != 'a' {
			if i < n/2 {
				flag = true
				ans[i] = 'a'
			}
			break
		}
	}
	if !flag {
		ans[n-1] = 'b'
	}

	return string(ans)
}
