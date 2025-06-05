package search_suggestions_system

import "sort"

type tries struct {
	children []*tries
	isEnd    bool
}

func (t *tries) insert(s string) {
	if len(s) == 0 {
		return
	}

	offset := s[0] - 'a'
	if t.children[offset] == nil {
		t.children[offset] = &tries{}
	}

	if len(s) == 1 {
		t.children[offset].isEnd = true
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	var query string
	iterLast := 0
	ans := make([][]string, 0)

	for _, ch := range searchWord {
		query += string(ch)
		iterFind := sort.Search(len(products)-iterLast, func(i int) bool {
			return products[i+iterLast] >= query
		}) + iterLast

		suggestion := make([]string, 0, 3)
		for i := iterFind; i < len(products) && len(suggestion) < 3; i++ {
			if len(products[i]) < len(query) {
				continue
			}

			if products[i][:len(query)] == query {
				suggestion = append(suggestion, products[i])
			}
		}

		ans = append(ans, suggestion)
		iterLast = iterFind
	}

	return ans
}
