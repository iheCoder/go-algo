package find_and_replace_in_string

import "sort"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	type pair struct {
		index          int
		source, target string
	}
	n := len(indices)
	ps := make([]*pair, 0, n)
	for i := 0; i < n; i++ {
		ps = append(ps, &pair{
			index:  indices[i],
			source: sources[i],
			target: targets[i],
		})
	}

	m := len(s)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].index < ps[j].index || (ps[i].index == ps[j].index && i+len(ps[i].source) <= m && s[ps[i].index:ps[i].index+len(ps[i].source)] == ps[i].source)
	})

	ans := make([]byte, 0, m)
	var j int
	for i := 0; i < m; {
		if j < n && ps[j].index == i && i+len(ps[j].source) <= m && s[i:i+len(ps[j].source)] == ps[j].source {
			ans = append(ans, []byte(ps[j].target)...)
			i += len(ps[j].source)
		} else {
			ans = append(ans, s[i])
			i++
		}

		for j < n && ps[j].index < i {
			j++
		}
	}

	return string(ans)
}
