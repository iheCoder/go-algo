package kth_smallest_element_in_a_sorted_matrix

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	// 检查小于等于mid数量是否大于等于k
	check := func(mid int) bool {
		i, j := n-1, 0
		var c int
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				// 换成通过i计算之后就好了，因为这是求小于等于的数量
				c += i + 1
				j++
			} else {
				i--
			}
		}
		return c >= k
	}

	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		mid := l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
