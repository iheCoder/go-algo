package online_stock_span

import "math"

type StockSpanner struct {
	items [][2]int
	idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{
		items: [][2]int{{-1, math.MaxInt32}},
		idx:   -1,
	}
}

func (s *StockSpanner) Next(price int) int {
	s.idx++
	for price >= s.items[len(s.items)-1][1] {
		s.items = s.items[:len(s.items)-1]
	}

	span := s.idx - s.items[len(s.items)-1][0]
	s.items = append(s.items, [2]int{s.idx, price})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
