package uncommon_words_from_two_sentences

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	m := make(map[string]int)
	insert := func(s string) {
		ss := strings.Split(s, " ")
		for _, x := range ss {
			m[x]++
		}
	}

	insert(s1)
	insert(s2)

	var ans []string
	for s, c := range m {
		if c == 1 {
			ans = append(ans, s)
		}
	}
	return ans
}
