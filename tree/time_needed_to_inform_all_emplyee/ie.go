package time_needed_to_inform_all_emplyee

type node struct {
	children   []*node
	informTime int
	trace      int
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	m := make([]node, n)
	for i, p := range manager {
		m[i].informTime = informTime[i]
		if p == -1 {
			continue
		}
		m[p].children = append(m[p].children, &m[i])
	}

	stk := []*node{&m[headID]}
	var ans int
	for len(stk) > 0 {
		p := stk[0]
		stk = stk[1:]

		for _, child := range p.children {
			child.trace = p.trace + p.informTime
			ans = maxInt(ans, child.trace)

			if len(child.children) > 0 {
				stk = append(stk, child)
			}
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
