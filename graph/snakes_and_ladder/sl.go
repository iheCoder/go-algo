package snakes_and_ladder

func id2rc(id, n int) (int, int) {
	r, c := (id-1)/n, (id-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return r, c
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	type pair struct {
		id, step int
	}
	q := []pair{
		{
			id:   1,
			step: 0,
		},
	}
	visited := make([]bool, n*n+1)
	end := n * n

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for step := 1; step <= 6; step++ {
			next := p.id + step
			if next > end {
				break
			}

			r, c := id2rc(next, n)
			if board[r][c] > 0 {
				next = board[r][c]
			}

			if next == end {
				return p.step + 1
			}

			if !visited[next] {
				visited[next] = true
				q = append(q, pair{
					id:   next,
					step: p.step + 1,
				})
			}
		}
	}

	return -1
}
