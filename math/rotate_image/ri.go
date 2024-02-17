package rotate_image

func rotate(matrix [][]int) {
	if len(matrix) <= 1 {
		return
	}

	n := len(matrix)
	var status int
	var i, j int
	x := matrix[i][j]
	for {
		switch {
		case status < 4:
			i, j = j, n-1-i
			x, matrix[i][j] = matrix[i][j], x
			status++
		default:
			status = 0
			if j < n-1-i-1 {
				j++
				x = matrix[i][j]
				continue
			}
			if i < n>>1-1 {
				i, j = i+1, i+1
				x = matrix[i][j]
				continue
			}
			return
		}
	}
}
