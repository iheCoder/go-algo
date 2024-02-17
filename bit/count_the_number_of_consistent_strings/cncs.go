package count_the_number_of_consistent_strings

func countConsistentStrings(allowed string, words []string) int {
	m := make([]bool, 26)
	for _, v := range allowed {
		m[v-'a'] = true
	}
	var c int
	var repeat bool
	for _, word := range words {
		repeat = false
		for _, w := range word {
			if !m[w-'a'] {
				repeat = true
				break
			}
		}
		if !repeat {
			c++
		}
	}
	return c
}
