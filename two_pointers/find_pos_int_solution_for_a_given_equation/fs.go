package find_pos_int_solution_for_a_given_equation

func findSolution(customFunction func(int, int) int, z int) [][]int {
	low, high1 := 1, 1000
	var mid int
	for low < high1 {
		mid = low + (high1-low)/2
		if customFunction(1, mid) < z {
			low = mid + 1
		} else {
			high1 = mid
		}
	}

	low = 1
	high2 := 1000
	for low < high2 {
		mid = low + (high2-low)/2
		if customFunction(mid, 1) < z {
			low = mid + 1
		} else {
			high2 = mid
		}
	}
	high := high1
	if high2 > high {
		high = high2
	}

	r := make([][]int, 0)
	i, j := 1, high
	for i <= high && j >= 1 {
		if customFunction(i, j) > z {
			j--
		} else if customFunction(i, j) < z {
			i++
		} else {
			r = append(r, []int{i, j})
			i++
			j--
		}
	}

	return r
}
