package max_len_of_pair_chain

import "sort"

func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	if n == 0 {
		return 0
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	ans := 1
	last := pairs[0][1]
	for i := 1; i < n; i++ {
		if pairs[i][0] > last {
			last = pairs[i][1]
			ans++
		}
	}

	return ans
}
