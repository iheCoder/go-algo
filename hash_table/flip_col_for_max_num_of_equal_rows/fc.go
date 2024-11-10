package flip_col_for_max_num_of_equal_rows

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	mp := make(map[string]int)
	for i := 0; i < m; i++ {
		arr := make([]byte, n)
		for j := 0; j < n; j++ {
			if matrix[i][j]^matrix[i][0] == 0 {
				arr[j] = '0'
			} else {
				arr[j] = '1'
			}
		}
		mp[string(arr)]++
	}

	var res int
	for _, val := range mp {
		if val > res {
			res = val
		}
	}
	return res
}
