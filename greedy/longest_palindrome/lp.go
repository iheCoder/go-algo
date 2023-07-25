package longest_palindrome

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	var total int
	var existOdd bool
	for _, c := range m {
		if c&1 > 0 {
			existOdd = true
			total += c - 1
		} else {
			total += c
		}
	}
	if existOdd {
		total++
	}
	return total
}

// 直接求偶数和+最大奇数是不对的，因为不一定需要把所有字符都用到
