package longest_string_chain

import "sort"

// 这个问题实在是太麻烦了
func longestStrChainDep(words []string) int {
	if len(words) <= 1 {
		return 1
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) || (len(words[i]) == len(words[j]) && words[i] < words[j])
	})
	n := len(words[len(words)-1])
	if len(words[0]) == n {
		return 1
	}
	ls := make(map[int]int)
	for i := 1; i < len(words); i++ {
		if len(words[i]) != len(words[i-1]) {
			ls[len(words[i-1])] = i - 1
		}
	}
	ls[n] = n - 1

	r := 1
	el := len(words) - 1
	sl := el
	var i, j int
	var d int
	var ok bool
	var ss, es string
	var si, ei int
	token := 1
	cr := 1
	for {
		for el >= 0 && d < r {
			if _, ok = ls[el]; !ok {
				el--
				sl = el
				d = 0
				continue
			}
			d++
		}

		si = ls[sl]
		ei = ls[el]
		ss = words[si]
		es = words[ei]

		i, j = 0, 0
		for i < len(ss) && j < len(es) {
			if es[j] != ss[i] {
				if token > 0 {
					token--
					j++
					continue
				}
				si--
				if si < 0 || len(words[si]) < len(words[si+1]) {
					ei--
					if ei < 0 || len(words[ei]) < len(words[ei+1]) {
						if cr > r {
							r = cr
						}
						cr = 1
						break
					}
					continue
				}
			}
			i++
			j++
		}
		cr++
		ei = si
		//si =

	}

	return r
}
