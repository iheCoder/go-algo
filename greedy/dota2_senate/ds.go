package dota2_senate

type senator struct {
	party   int
	blocked bool
}

var m = map[int]string{
	1: "Dire",
	2: "Radiant",
}

func predictPartyVictory(senate string) string {
	n := len(senate)
	ss := make([]*senator, n)
	var dc, rc int
	for i := 0; i < n; i++ {
		switch senate[i] {
		case 'D':
			ss[i] = &senator{party: 1}
			dc++
		case 'R':
			ss[i] = &senator{party: 2}
			rc++
		}
	}

	for i := 0; ; i++ {
		if i == n {
			i = 0
		}
		if ss[i].blocked {
			continue
		}
		switch ss[i].party {
		case 1:
			if rc <= 0 {
				return m[1]
			}
			for j := 1; j < n; j++ {
				k := (i + j) % n
				if ss[k].party != 1 && !ss[k].blocked {
					ss[k].blocked = true
					rc--
					break
				}
			}
		case 2:
			if dc <= 0 {
				return m[2]
			}
			for j := 1; j < n; j++ {
				k := (i + j) % n
				if ss[k].party != 2 && !ss[k].blocked {
					ss[k].blocked = true
					dc--
					break
				}
			}
		}
	}
}
