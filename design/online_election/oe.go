package online_election

import "sort"

type voteSituation struct {
	t, whoWin int
}

type TopVotedCandidate struct {
	vs []*voteSituation
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	vs := make([]*voteSituation, 0, n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		vs = append(vs, &voteSituation{
			t:      times[i],
			whoWin: persons[i],
		})
		m[persons[i]] = 0
	}

	sort.Slice(vs, func(i, j int) bool {
		return vs[i].t < vs[j].t
	})

	mx := -1
	for i := 0; i < n; i++ {
		m[vs[i].whoWin]++
		if mx == -1 || m[vs[i].whoWin] >= m[mx] {
			mx = vs[i].whoWin
		} else {
			vs[i].whoWin = mx
		}
	}

	return TopVotedCandidate{
		vs: vs,
	}
}

func (c *TopVotedCandidate) Q(t int) int {
	n := len(c.vs)
	x := sort.Search(n, func(i int) bool {
		return c.vs[i].t >= t
	})
	if x == n || c.vs[x].t != t {
		x--
	}

	return c.vs[x].whoWin
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
