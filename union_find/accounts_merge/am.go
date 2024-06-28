package accounts_merge

import "sort"

type unionFind struct {
	parents []int
}

func newUnionFind(size int) *unionFind {
	parents := make([]int, size)
	for i := 0; i < size; i++ {
		parents[i] = i
	}

	return &unionFind{
		parents: parents,
	}
}

func (u *unionFind) find(x int) int {
	if u.parents[x] != x {
		p := u.find(u.parents[x])
		u.parents[x] = p
	}

	return u.parents[x]
}

func (u *unionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}

	u.parents[yp] = xp
}

func accountsMerge(accounts [][]string) [][]string {
	emailIndex := make(map[string]int)
	emailName := make(map[string]string)
	for _, account := range accounts {
		name := account[0]
		for _, s := range account[1:] {
			if _, has := emailIndex[s]; !has {
				emailIndex[s] = len(emailIndex)
				emailName[s] = name
			}
		}
	}

	uf := newUnionFind(len(emailIndex))
	for _, account := range accounts {
		firstIndex := account[1]
		for _, s := range account[2:] {
			uf.union(emailIndex[firstIndex], emailIndex[s])
		}
	}

	em := make(map[int][]string)
	for s, i := range emailIndex {
		pi := uf.find(i)
		em[pi] = append(em[pi], s)
	}

	var ans [][]string
	for _, emails := range em {
		sort.Strings(emails)
		ans = append(ans, append([]string{emailName[emails[0]]}, emails...))
	}
	return ans
}
