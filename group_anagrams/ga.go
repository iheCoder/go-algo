package group_anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		k := sortString(str)
		if m[k] == nil {
			m[k] = []string{str}
		} else {
			m[k] = append(m[k], str)
		}
	}

	r := make([][]string, 0)
	for _, ss := range m {
		r = append(r, ss)
	}
	return r
}

func sortString(s string) string {
	sb := []byte(s)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})

	return string(sb)
}
