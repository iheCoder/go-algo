package xor_queries_of_a_subarray

func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	xors := make([]int, n+1)
	for i := 0; i < n; i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}

	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		ans = append(ans, xors[query[0]]^xors[query[1]+1])
	}

	return ans
}
