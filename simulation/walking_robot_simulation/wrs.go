package walking_robot_simulation

import "sort"

type direction int

const (
	north direction = iota
	south
	west
	east
)

func robotSim(commands []int, obstacles [][]int) int {
	if len(commands) == 0 {
		return 0
	}

	rowm := make(map[int][]int)
	colm := make(map[int][]int)

	for _, obstacle := range obstacles {
		x, y := obstacle[0], obstacle[1]
		rowm[x] = append(rowm[x], y)
		i := len(rowm[x]) - 1
		for i-1 >= 0 && rowm[x][i-1] > rowm[x][i] {
			rowm[x][i-1], rowm[x][i] = rowm[x][i], rowm[x][i-1]
			i--
		}
		colm[y] = append(colm[y], x)
		i = len(colm[y]) - 1
		for i-1 >= 0 && colm[y][i-1] > colm[y][i] {
			colm[y][i-1], colm[y][i] = colm[y][i], colm[y][i-1]
			i--
		}
	}

	d := north
	x, y := 0, 0
	var ans int
	var walk func(step int)
	walk = func(step int) {
		switch d {
		case north:
			obs, ok := rowm[x]
			if ok {
				oy := sort.SearchInts(obs, y)
				if oy < len(obs) && obs[oy] != y && obs[oy] <= y+step {
					y = obs[oy] - 1
					return
				}
			}
			y = y + step

		case south:
			obs, ok := rowm[x]
			if ok {
				oy := sort.SearchInts(obs, y) - 1
				if oy >= 0 && obs[oy] != y && obs[oy] >= y-step {
					y = obs[oy] + 1
					return
				}
			}
			y = y - step

		case east:
			obs, ok := colm[y]
			if ok {
				ox := sort.SearchInts(obs, x)
				if ox < len(obs) && obs[ox] != x && obs[ox] <= x+step {
					x = obs[ox] - 1
					return
				}
			}
			x = x + step

		case west:
			obs, ok := colm[y]
			if ok {
				ox := sort.SearchInts(obs, x) - 1
				if ox >= 0 && obs[ox] != x && obs[ox] >= x-step {
					x = obs[ox] + 1
					return
				}
			}
			x = x - step

		}
	}
	for _, command := range commands {
		switch command {
		case -2:
			d = turnLeft(d)
		case -1:
			d = turnRight(d)
		default:
			walk(command)
			cs := x*x + y*y
			if cs > ans {
				ans = cs
			}
		}
	}
	return ans
}

func turnLeft(d direction) direction {
	switch d {
	case north:
		return west
	case west:
		return south
	case south:
		return east
	case east:
		return north
	}
	panic("wrong direction")
}

func turnRight(d direction) direction {
	switch d {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}
	panic("wrong direction")
}
