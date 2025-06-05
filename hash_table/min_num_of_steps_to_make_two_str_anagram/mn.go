package min_num_of_steps_to_make_two_str_anagram

func minSteps(s string, t string) int {
	sm, tm := make(map[byte]int), make(map[byte]int)
	for i := range s {
		sm[s[i]]++
	}
	for i := range t {
		tm[t[i]]++
	}

	var ans int
	for b, c := range sm {
		d := c - tm[b]
		ans += abs(d)
		delete(tm, b)
	}

	for _, c := range tm {
		ans += c
	}

	return ans / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
