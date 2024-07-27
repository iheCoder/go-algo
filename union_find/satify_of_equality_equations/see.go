package satify_of_equality_equations

import "sort"

type unionFind struct {
	parents []uint8
}

func newUF() *unionFind {
	parents := make([]uint8, 26)
	for i := 0; i < 26; i++ {
		parents[i] = uint8(i)
	}
	return &unionFind{parents: parents}
}

func (u *unionFind) find(x uint8) uint8 {
	if u.parents[x] != x {
		u.parents[x] = u.find(u.parents[x])
	}
	return u.parents[x]
}

func (u *unionFind) union(x, y uint8) {
	rx, ry := u.find(x), u.find(y)
	if rx != ry {
		u.parents[rx] = ry
	}
}

func equationsPossible(equations []string) bool {
	sort.Slice(equations, func(i, j int) bool {
		if equations[i][1] == '=' {
			return true
		}
		return false
	})

	em := make([]uint8, 26)
	uf := newUF()
	for _, equation := range equations {
		x, y := equation[0]-'a', equation[3]-'a'
		if equation[1] == '=' {
			uf.union(x, y)
		} else {
			if em[x] > 0 && em[y] > 0 {
				rx, ry := uf.find(x), uf.find(y)
				if rx == ry {
					return false
				}
			} else if x == y {
				return false
			}
		}
		em[x]++
		em[y]++
	}
	return true
}
