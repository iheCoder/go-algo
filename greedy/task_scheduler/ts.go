package task_scheduler

import "sort"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	cnts := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		cnts[tasks[i]]++
	}

	me, mc := 0, 0
	for _, c := range cnts {
		if c > me {
			me, mc = c, 1
		} else if c == me {
			mc++
		}
	}

	t := (me-1)*(n+1) + mc
	if t < len(tasks) {
		return len(tasks)
	}
	return t
}

func leastIntervalFetchMax(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	m := make(map[byte]int)
	s := make([]byte, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
		if m[tasks[i]] == 1 {
			s = append(s, tasks[i])
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})

	var ans int
	used := n
	for i := 0; i < len(s); i++ {
		for m[s[i]] > 0 {
			ans += n - used
			ans++
			m[s[i]]--
			used = 0
			for j := 0; used < n && i+j+1 < len(s); j++ {
				if m[s[i+j+1]] <= 0 {
					continue
				}
				m[s[i+j+1]]--
				ans++
				used++
			}
		}
	}
	return ans
}
