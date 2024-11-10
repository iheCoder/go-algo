package swap_for_longest_repeated_char_substring

func maxRepOpt1(text string) int {
	m := make(map[byte]int)
	n := len(text)
	for i := 0; i < n; i++ {
		m[text[i]]++
	}

	var i, ans int
	for i < n {
		win, j := 1, i+1
		for j < n && text[i] == text[j] {
			win++
			j++
		}

		if win < m[text[i]] {
			ans = maxInt(ans, win+1)
		} else {
			ans = maxInt(ans, win)
		}

		if j < n-1 && text[j] != text[i] && text[i] == text[j+1] {
			extWin, k := win+1, j+1
			for k < n && text[k] == text[i] {
				extWin++
				k++
			}

			if extWin <= m[text[i]] {
				ans = maxInt(ans, extWin)
			} else {
				ans = maxInt(ans, extWin-1)
			}
		}

		i = j
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
