package image_overlap

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	l := 2*n + 1
	m := make([][]int, l)
	for i := 0; i < l; i++ {
		m[i] = make([]int, l)
	}

	for i1 := 0; i1 < n; i1++ {
		for j1 := 0; j1 < n; j1++ {
			if img1[i1][j1] == 1 {
				for i2 := 0; i2 < n; i2++ {
					for j2 := 0; j2 < n; j2++ {
						if img2[i2][j2] == 1 {
							m[i1-i2+n][j1-j2+n]++
						}
					}
				}
			}
		}
	}

	var ans int
	for _, mi := range m {
		for _, x := range mi {
			if x > ans {
				ans = x
			}
		}
	}
	return ans
}
