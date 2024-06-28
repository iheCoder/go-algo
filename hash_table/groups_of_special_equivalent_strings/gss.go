package groups_of_special_equivalent_strings

import "sort"

func numSpecialEquivGroups(words []string) int {
	n := len(words[0])
	m := make(map[string]bool)
	for _, word := range words {
		l, r := make([]byte, 0, n/2), make([]byte, 0, n/2)
		for i := 0; i < n; i++ {
			if i%2 == 1 {
				l = append(l, word[i])
			} else {
				r = append(r, word[i])
			}
		}

		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})

		m[string(l)+"_"+string(r)] = true
	}

	return len(m)
}
