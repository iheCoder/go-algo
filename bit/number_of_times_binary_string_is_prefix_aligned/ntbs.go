package number_of_times_binary_string_is_prefix_aligned

func numTimesAllBlue(flips []int) int {
	ans, n, right := 0, len(flips), 0
	for i := 0; i < n; i++ {
		if flips[i] > right {
			right = flips[i]
		}
		if right == i+1 {
			ans++
		}
	}

	return ans
}
