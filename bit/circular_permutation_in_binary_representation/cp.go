package circular_permutation_in_binary_representation

func circularPermutation(n int, start int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = (i >> 1) ^ i ^ start
	}

	return ans
}
