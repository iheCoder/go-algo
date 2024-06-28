package exclusive_time_of_functions

import (
	"strconv"
	"strings"
)

type processLog struct {
	isStart      bool
	id           int
	curTime      int
	innerSumCost int
}

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	stk := make([]*processLog, 0)
	for _, l := range logs {
		pl := parseIntoLog(l)
		if len(stk) > 0 {
			last := stk[len(stk)-1]
			if checkMatch(last, pl) {
				cur := pl.curTime - last.curTime + 1 - last.innerSumCost
				ans[pl.id] += cur
				stk = stk[:len(stk)-1]
				if len(stk) > 0 {
					stk[len(stk)-1].innerSumCost += cur + last.innerSumCost
				}
				continue
			}
		}

		stk = append(stk, pl)
	}

	return ans
}

func parseIntoLog(l string) *processLog {
	a := strings.Split(l, ":")
	id, _ := strconv.Atoi(a[0])
	isStart := a[1] == "start"
	st, _ := strconv.Atoi(a[2])
	return &processLog{
		isStart: isStart,
		id:      id,
		curTime: st,
	}
}

func checkMatch(pl1, pl2 *processLog) bool {
	return pl1.id == pl2.id && pl1.isStart != pl2.isStart
}
