package assign_cookies

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	var i, j int
	var r int
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			r++
			i++
		}
		j++
	}
	return r
}
