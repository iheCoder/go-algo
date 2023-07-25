package spiral_matrix_ii

type direction int

const (
	right direction = iota
	down
	left
	top
)

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	d := right
	x, y := -1, -1
	for c := 1; c <= n*n; {
		switch d {
		case right:
			x++
			y++
			for y < n && m[x][y] == 0 {
				m[x][y] = c
				c++
				y++
			}
		case down:
			y--
			x++
			for x < n && m[x][y] == 0 {
				m[x][y] = c
				c++
				x++
			}
		case left:
			x--
			y--
			for y >= 0 && m[x][y] == 0 {
				m[x][y] = c
				c++
				y--
			}
		case top:
			y++
			x--
			for x >= 0 && m[x][y] == 0 {
				m[x][y] = c
				c++
				x--
			}
		}
		d = turnDirection(d)
	}
	return m
}

func turnDirection(d direction) direction {
	switch d {
	case right:
		return down
	case down:
		return left
	case left:
		return top
	case top:
		return right
	}
	panic("wrong direction")
}
