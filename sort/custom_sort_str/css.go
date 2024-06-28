package custom_sort_str

import "sort"

func customSortString(order string, s string) string {
	ho := make([]int, 26)
	var w int
	for i := 0; i < len(order); i++ {
		ho[order[i]-'a'] = w
		w++
	}

	sb := make([]byte, len(s))
	copy(sb, s)
	sort.Slice(sb, func(i, j int) bool {
		return ho[sb[i]-'a'] < ho[sb[j]-'a']
	})
	return string(sb)
}
