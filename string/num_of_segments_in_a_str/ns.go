package num_of_segments_in_a_str

func countSegments(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	lastSpace := true
	var count int
	for i := 0; i < n; i++ {
		if lastSpace && s[i] != ' ' {
			count++
		}
		lastSpace = s[i] == ' '
	}

	return count
}
