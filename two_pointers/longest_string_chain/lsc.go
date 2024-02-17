package longest_string_chain

import "sort"

func longestStrChain(words []string) int {
	if len(words) <= 1 {
		return 1
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	if len(words[0]) == len(words[len(words)-1]) {
		return 1
	}

	d := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		d[i] = 1
	}

	r := 1
	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if isMatch(words[j], words[i]) {
				d[i] = maxInt(d[i], d[j]+1)
			}
		}
		r = maxInt(r, d[i])
	}
	return r
}

func isMatch(a, b string) bool {
	if len(a) != len(b)-1 {
		return false
	}

	var i int
	for j := 0; j < len(b) && i < len(a); j++ {
		if a[i] == b[j] {
			i++
		}
	}

	return i == len(a)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
