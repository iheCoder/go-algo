package max_xor_of_two_numbers_in_an_array

func findMaximumXOR(nums []int) int {
	highBit := 30
	var x int
	for i := highBit; i >= 0; i-- {
		seen := make(map[int]bool)
		for _, num := range nums {
			seen[num>>i] = true
		}

		xn := x<<1 | 1
		found := false
		for _, num := range nums {
			if seen[num>>i^xn] {
				found = true
				break
			}
		}

		if found {
			x = xn
		} else {
			x = xn - 1
		}
	}

	return x
}
