package reveal_cards_in_increasing_order

import (
	"container/list"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	index := list.New()
	n := len(deck)
	for i := 0; i < n; i++ {
		index.PushBack(i)
	}

	res := make([]int, n)
	for _, card := range deck {
		i := index.Remove(index.Front()).(int)
		res[i] = card
		if index.Len() > 0 {
			index.PushBack(index.Remove(index.Front()))
		}
	}

	return res
}

// 2 13 3 11 5 17 7
// 3 11 5 17 7 13
// 5 17 7 13 11
// 7 13 11 17
// 11 17 13
// 13 17
// 17
