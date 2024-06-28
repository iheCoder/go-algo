package prison_cells_after_n_days

import "fmt"

// 1 0 1
// 0 1 0
// 0 1 0

// 1 1 1
// 0 1 0

// 1 0 1 0
// 0 1 1 0
// 0 0 0 0
// 0 1 1 0
func prisonAfterNDays(cells []int, n int) []int {
	nextDay := func(cs []int) []int {
		next := make([]int, 8)
		for i := 1; i < 7; i++ {
			if cs[i-1] == cs[i+1] {
				next[i] = 1
			} else {
				next[i] = 0
			}
		}

		return next
	}

	seen := make(map[string]int)
	for n > 0 {
		key := fmt.Sprintf("%v", cells)
		if day, ok := seen[key]; ok {
			n %= day - n
		}

		if n > 0 {
			seen[key] = n
			n--
			cells = nextDay(cells)
		}
	}

	return cells
}
