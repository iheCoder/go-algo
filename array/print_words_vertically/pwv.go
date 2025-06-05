package print_words_vertically

import "strings"

func printVertically(s string) []string {
	n := len(s)
	matrix := make([][]byte, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]byte, n)
	}

	var ml, row, col int
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			ml = maxInt(ml, col+1)
			row++
			col = 0
			continue
		}

		matrix[row][col] = s[i]
		col++
	}
	ml = maxInt(ml, col)

	var ans []string
	for i := 0; i < ml; i++ {
		temp := make([]byte, 0)
		for j := 0; j <= row; j++ {
			if matrix[j][i] == 0 {
				matrix[j][i] = ' '
			}
			temp = append(temp, matrix[j][i])
		}
		ans = append(ans, strings.TrimRight(string(temp), " "))
	}

	for len(ans) > 0 && len(ans[len(ans)-1]) == 0 {
		ans = ans[:len(ans)-1]
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
