package find_the_longest_substring_containing_vowels_in_even_counts

func findTheLongestSubstring(s string) int {
	ans, status := 0, 0
	pos := make([]int, 1<<5)
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	pos[0] = 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}

		if pos[status] >= 0 {
			ans = maxInt(ans, i+1-pos[status])
		} else {
			pos[status] = i + 1
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
