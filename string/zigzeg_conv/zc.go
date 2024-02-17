package zigzeg_conv

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	var j int
	var up bool
	for i := 0; i < len(s); i++ {
		matrix[j] = append(matrix[j], s[i])
		if up {
			j--
		} else {
			j++
		}
		if j == 0 || j == numRows-1 {
			up = !up
		}
	}

	ans := make([]byte, 0, numRows)
	for _, m := range matrix {
		ans = append(ans, m...)
	}
	return string(ans)
}
