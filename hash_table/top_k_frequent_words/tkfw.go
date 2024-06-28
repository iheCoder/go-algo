package top_k_frequent_words

import "sort"

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	ws := make([]string, 0, len(m))
	for s, _ := range m {
		ws = append(ws, s)
	}

	sort.Slice(ws, func(i, j int) bool {
		return m[ws[i]] > m[ws[j]] || (m[ws[i]] == m[ws[j]] && ws[i] < ws[j])
	})

	return ws[:k]
}
