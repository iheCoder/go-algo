package magic_squares_in_grid

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}

	var ans int
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			if grid[i+1][j+1] != 5 {
				continue
			}
			rm := make(map[int]int)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					rm[grid[i+k][j+l]]++
				}
			}
			var repeat bool
			for x, c := range rm {
				if c > 1 || x < 1 || x > 9 {
					repeat = true
					break
				}
			}
			if repeat {
				continue
			}

			var total int
			total = grid[i][j] + grid[i][j+1] + grid[i][j+2]
			if grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] != total {
				continue
			}
			if grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] != total {
				continue
			}

			if grid[i][j]+grid[i+1][j]+grid[i+2][j] != total {
				continue
			}
			if grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] != total {
				continue
			}
			if grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] != total {
				continue
			}

			if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != total {
				continue
			}
			if grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] != total {
				continue
			}
			ans++
		}
	}
	return ans
}
