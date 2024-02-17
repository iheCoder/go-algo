package min_number_of_swaps_to_make_the_string_balanced

func minSwaps(s string) int {
	balanced := 0
	var i, j int
	for {
		for ; i < len(s) && s[i] != '['; i++ {
		}
		if i == len(s) {
			break
		}
		for j = maxInt(i+1, j); j < len(s) && s[j] != ']'; j++ {
		}
		if j == len(s) {
			break
		}
		balanced += 2
		i++
		j++
	}
	if balanced == len(s) {
		return 0
	}

	return ((len(s) - balanced - 1) >> 2) + 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
