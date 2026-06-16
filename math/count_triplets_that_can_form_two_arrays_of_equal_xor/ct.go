package count_triplets_that_can_form_two_arrays_of_equal_xor

func countTriplets(arr []int) int {
	seen := make(map[int][]int)
	seen[0] = []int{-1}
	var sum, ans int

	for i, item := range arr {
		sum ^= item
		for j := 0; j < len(seen[sum]); j++ {
			ans += i - seen[sum][j] - 1
		}
		seen[sum] = append(seen[sum], i)
	}

	return ans
}
