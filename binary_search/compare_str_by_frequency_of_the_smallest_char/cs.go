package compare_str_by_frequency_of_the_smallest_char

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	getSF := func(s string) int {
		m := byte('z')
		for i := 0; i < len(s); i++ {
			if s[i] < m {
				m = s[i]
			}
		}

		var c int
		for i := 0; i < len(s); i++ {
			if s[i] == m {
				c++
			}
		}

		return c
	}

	wsf := make([]int, 0, len(words))
	for _, word := range words {
		wsf = append(wsf, getSF(word))
	}
	sort.Ints(wsf)

	ans := make([]int, 0, len(queries))
	n := len(wsf)
	for _, query := range queries {
		target := getSF(query)

		x := sort.Search(n, func(i int) bool {
			return wsf[i] > target
		})
		y := n - x
		ans = append(ans, y)
	}

	return ans
}
