package range_sum_query_2d_immutable

type NumMatrix struct {
	sm [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sm := make([][]int, m)
	for i := 0; i < m; i++ {
		sm[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		sm[i][0] = matrix[i][0]
		if i > 0 {
			sm[i][0] += sm[i-1][0]
		}
	}
	for i := 0; i < n; i++ {
		sm[0][i] = matrix[0][i]
		if i > 0 {
			sm[0][i] += sm[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sm[i][j] = sm[i-1][j] + sm[i][j-1] + matrix[i][j] - sm[i-1][j-1]
		}
	}

	return NumMatrix{sm: sm}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	r := this.sm[row2][col2]
	if row1 > 0 {
		r -= this.sm[row1-1][col2]
	}
	if col1 > 0 {
		r -= this.sm[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		r += this.sm[row1-1][col1-1]
	}
	return r
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
