package count_number_teams

import "sort"

func numTeamsBIT(rating []int) int {
	n := len(rating)

	sorted := make([]int, n)
	copy(sorted, rating)
	sort.Ints(sorted)

	getID := func(x int) int {
		return sort.SearchInts(sorted, x) + 1
	}

	leftLess := make([]int, n)
	leftGreater := make([]int, n)
	rightLess := make([]int, n)
	rightGreater := make([]int, n)

	bt := NewBIT(n)

	for i := 0; i < n; i++ {

	}
}

type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+1), // 1-based
		n:    n,
	}
}

func (b *BIT) lowbit(x int) int {
	return x & -x
}

func (b *BIT) Add(idx, delta int) {
	for idx <= b.n {
		b.tree[idx] += delta
		idx += b.lowbit(idx)
	}
}

func (b *BIT) Sum(idx int) int {
	var res int
	for idx > 0 {
		res += b.tree[idx]
		idx -= b.lowbit(idx)
	}

	return res
}
