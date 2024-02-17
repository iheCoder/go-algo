package non_overlapping_intervals

import "sort"

type graphItem struct {
	index   int
	depends map[int]bool
	c       int
}

func NewGraphItem(index int) *graphItem {
	return &graphItem{index: index}
}

func (g *graphItem) addDepend(w int) {
	g.depends[w] = true
	g.c++
}

type graph struct {
	items []*graphItem
	c     int
}

func NewGraph(items []*graphItem, c int) *graph {
	sort.Slice(items, func(i, j int) bool {
		return items[i].c > items[j].c
	})
	return &graph{
		items: items,
		c:     c,
	}
}

func (g *graph) pop() *graphItem {
	x := g.items[0]
	g.c -= x.c
	x.c = 0
	g.items = g.items[1:]
	if len(g.items) == 0 {
		return x
	}
	for i := len(g.items) - 1; i >= 0; i-- {
		if !x.depends[g.items[i].index] {
			continue
		}
		g.items[i].c--
		g.c--
		j := i
		for j+1 < len(g.items) && g.items[j].c < g.items[j+1].c {
			g.items[j], g.items[j+1] = g.items[j+1], g.items[j]
			j++
		}
	}

	return x
}

func (g *graph) existDepend() bool {
	return g.c > 0
}

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})

	right := intervals[0][1]
	num := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			num++
			right = intervals[i][1]
		}
	}

	return n - num
}
