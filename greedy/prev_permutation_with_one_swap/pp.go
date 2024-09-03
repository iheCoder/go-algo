package prev_permutation_with_one_swap

func prevPermOpt1(arr []int) []int {
	n := len(arr)
	sp := n - 2
	for sp >= 0 {
		for i := n - 1; i > sp; i-- {
			if arr[i] < arr[sp] && (i-1 <= sp || arr[i] != arr[i-1]) {
				arr[i], arr[sp] = arr[sp], arr[i]
				return arr
			}
		}
		sp--
	}

	return arr
}
