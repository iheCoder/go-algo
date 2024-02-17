package sort_char_by_frequency

import "sort"

// 与其插入排序，不如直接hash之后再排序
type pair struct {
	count int
	b     byte
}

type hashPairs struct {
	ps []*pair
	h  []int
}

func newHashPairs() *hashPairs {
	h := make([]int, 52)
	for i := 0; i < 52; i++ {
		h[i] = -1
	}
	return &hashPairs{
		ps: make([]*pair, 0),
		h:  h,
	}
}

func (h *hashPairs) insert(b byte) {
	var diff int
	if b >= 'a' {
		diff = int(b-'a') + 26
	} else {
		diff = int(b - 'A')
	}
	if h.h[diff] < 0 {
		h.ps = append(h.ps, &pair{
			count: 0,
			b:     b,
		})
		h.h[diff] = len(h.ps) - 1
	}
	i := h.h[diff]
	h.ps[i].count++
	for i > 0 && h.ps[i-1].count < h.ps[i].count {
		h.ps[i], h.ps[i-1] = h.ps[i-1], h.ps[i]
		i--
	}
}

func frequencySort(s string) string {
	h := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		h[s[i]]++
	}

	sc := make([]*pair, 0, len(h))
	for b, c := range h {
		sc = append(sc, &pair{
			count: c,
			b:     b,
		})
	}

	sort.Slice(sc, func(i, j int) bool {
		return sc[i].count > sc[j].count
	})

	ans := make([]byte, 0, len(s))
	for _, p := range sc {
		for i := 0; i < p.count; i++ {
			ans = append(ans, p.b)
		}
	}

	return string(ans)
}
