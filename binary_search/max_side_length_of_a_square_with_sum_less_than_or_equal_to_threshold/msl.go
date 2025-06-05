package max_side_length_of_a_square_with_sum_less_than_or_equal_to_threshold

func getRect(P [][]int, x1, y1, x2, y2 int) int {
	return P[x2][y2] - P[x1-1][y2] - P[x2][y1-1] + P[x1-1][y1-1]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	p := make([][]int, m+1)
	for i := range p {
		p[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			p[i][j] = p[i-1][j] + p[i][j-1] - p[i-1][j-1] + mat[i-1][j-1]
		}
	}

	r, ans := minInt(m, n), 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for c := ans + 1; c <= r; c++ {
				if i+c-1 <= m && j+c-1 <= n && getRect(p, i, j, i+c-1, j+c-1) <= threshold {
					ans++
				} else {
					break
				}
			}
		}
	}

	return ans
}

//func maxSideLength(mat [][]int, threshold int) int {
//	m, n := len(mat), len(mat[0])
//	P := make([][]int, m+1)
//	for i := range P {
//		P[i] = make([]int, n+1)
//	}
//
//	for i := 1; i <= m; i++ {
//		for j := 1; j <= n; j++ {
//			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + mat[i-1][j-1]
//		}
//	}
//
//	l, r, ans := 1, minInt(m, n), 0
//	for l < r {
//		mid := l + (r-l)/2
//		find := false
//		for i := 1; i <= m-mid+1; i++ {
//			for j := 1; j <= n-mid+1; j++ {
//				if getRect(P, i, j, i+mid-1, j+mid-1) <= threshold {
//					find = true
//				}
//			}
//		}
//
//		if find {
//			ans = mid
//			l = mid + 1
//		} else {
//			r = mid - 1
//		}
//	}
//
//	return ans
//}
