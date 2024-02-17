package jump_game_4

func canReach(s string, minJump int, maxJump int) bool {
	if s[len(s)-1] != '0' {
		return false
	}

	rs := make([]int, len(s))
	sums := make([]int, len(s))
	// 既然是前缀和，那么index为0的一定能到达
	for i := 0; i < len(sums); i++ {
		sums[i] = 1
	}

	var total int
	var l, r int
	for i := minJump; i < len(s); i++ {
		l, r = i-maxJump, i-minJump
		if s[i] == '0' {
			if l <= 0 {
				total = sums[r]
			} else {
				total = sums[r] - sums[l-1]
			}
			if total > 0 {
				rs[i] = 1
			}
		}
		sums[i] = sums[i-1] + rs[i]
	}

	return rs[len(s)-1] > 0
}
