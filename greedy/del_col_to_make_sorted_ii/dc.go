package del_col_to_make_sorted_ii

func minDeletionSize(strs []string) int {
	m, n := len(strs), len(strs[0])
	cuts := make([]bool, m-1)
	var ans int
	cols := make([][]byte, n)
	for i := range cols {
		cols[i] = make([]byte, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cols[j][i] = strs[i][j]
		}
	}

	for _, col := range cols {
		valid := true
		for i := 0; i < len(col)-1; i++ {
			if !cuts[i] && col[i] > col[i+1] {
				valid = false
				ans++
				break
			}
		}

		if valid {
			for i := 0; i < len(col)-1; i++ {
				if col[i] < col[i+1] {
					cuts[i] = true
				}
			}
		}
	}

	return ans
}

// 	m := len(strs)
//	n := len(strs[0])
//
//	var res int
//	var sorted bool
//	for i := 0; i < n; i++ {
//		sorted = true
//		for j := 1; j < m; j++ {
//			if strs[j][i] < strs[j-1][i] {
//				res++
//				sorted = false
//				break
//			}
//		}
//
//		if sorted {
//			return res
//		}
//	}
//
//	return res
