package min_domino_rotations_for_equal_row

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	cs := make([]int, 6)

	for _, top := range tops {
		cs[top-1]++
	}
	for _, bottom := range bottoms {
		cs[bottom-1]++
	}

	target := -1
	for i, c := range cs {
		if c >= n {
			target = i + 1
			break
		}
	}
	if target == -1 {
		return -1
	}

	var count int
	pc, nc := 0, 0
	for i, top := range tops {
		if top == target || bottoms[i] == target {
			count++
		}
		if top == target && target != bottoms[i] {
			nc++
		} else if top != target && target == bottoms[i] {
			pc++
		}
	}
	if count < n {
		return -1
	}

	if nc < pc {
		return nc
	}
	return pc
}
