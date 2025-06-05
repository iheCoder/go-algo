package find_positive_int_solution_for_a_given_equation

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	limit := 1000
	for i := 1; i <= limit; i++ {
		for j := 1; j <= limit; j++ {
			if customFunction(i, j) == z {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}
