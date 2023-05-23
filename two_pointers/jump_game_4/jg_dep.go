package jump_game_4

func canReachTPDep(s string, minJump int, maxJump int) bool {
	if s[len(s)-1] != '0' {
		return false
	}

	curMinPos, curMaxPos := len(s)-1, len(s)-1
	nextMinPos, nextMaxPos := len(s)-1, len(s)-1
	var imin, imax int

	for {
		for i := curMinPos; i >= curMaxPos; i-- {
			if s[i] == '0' {
				imin = i - minJump
				imax = i - maxJump
				if imax <= 0 && imin >= 0 {
					return true
				}
				// 改成记录pos最大的nextMinPos是没用的，因为会走到一个根本去不了的地方
				nextMinPos = minInt(nextMinPos, imin)
				nextMaxPos = minInt(nextMaxPos, imax)
			}
		}
		if curMaxPos == nextMaxPos || nextMaxPos < 0 {
			return false
		}
		curMinPos = nextMinPos
		curMaxPos = nextMaxPos
		nextMinPos = -1
	}
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func canReachDP(s string, minJump int, maxJump int) bool {
	n := len(s)
	f, pre := make([]int, n), make([]int, n)
	f[0] = 1
	for i := 0; i < minJump; i++ {
		pre[i] = 1
	}
	var l, r int
	var total int
	for i := minJump; i < n; i++ {
		l, r = i-maxJump, i-minJump
		if s[i] == '0' {
			if l <= 0 {
				total = pre[r]
			} else {
				total = pre[r] - pre[l-1]
			}
			if total != 0 {
				f[i] = 1
			}
		}
		pre[i] = pre[i-1] + f[i]
	}

	return f[len(s)-1] > 0
}
