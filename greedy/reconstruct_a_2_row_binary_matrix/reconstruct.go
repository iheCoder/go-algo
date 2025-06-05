package reconstruct_a_2_row_binary_matrix

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	result := make([][]int, 2)
	result[0], result[1] = make([]int, len(colsum)), make([]int, len(colsum))

	for i := 0; i < len(colsum); i++ {
		switch colsum[i] {
		case 2:
			upper--
			lower--
			result[0][i], result[1][i] = 1, 1
		case 1:
			if upper >= lower {
				result[0][i] = 1
				upper--
			} else {
				result[1][i] = 1
				lower--
			}
		}

		if upper < 0 || lower < 0 {
			return nil
		}
	}

	if upper != 0 || lower != 0 {
		return nil
	}

	return result
}
