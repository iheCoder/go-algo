package escape_the_ghosts

import "math"

func escapeGhosts(ghosts [][]int, target []int) bool {
	p := abs(target[0]) + abs(target[1])
	mp := math.MaxInt
	for _, ghost := range ghosts {
		mp = minInt(mp, abs(target[0]-ghost[0])+abs(target[1]-ghost[1]))
	}
	return mp > p
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
