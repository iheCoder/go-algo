package min_genetic_mutation

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}

	sv, ed := genVal(startGene), genVal(endGene)
	m := make(map[int]bool)
	for _, s := range bank {
		m[genVal(s)] = false
	}
	if _, ok := m[ed]; !ok {
		return -1
	}

	var f func(count, cur int) int
	f = func(count, cur int) int {
		if cur == ed {
			return count
		}

		vn := -1
		for k, visited := range m {
			if !visited && checkOneChange(cur, k) {
				m[k] = true
				vn = minInt(vn, f(count+1, k))
				m[k] = false
			}
		}
		return vn
	}
	return f(0, sv)
}

func minInt(x, y int) int {
	if x == -1 {
		return y
	} else if y == -1 {
		return x
	}

	if x < y {
		return x
	}
	return y
}

func genVal(s string) int {
	var val int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			val = val*10 + 1
		case 'C':
			val = val*10 + 2
		case 'G':
			val = val*10 + 3
		case 'T':
			val = val*10 + 4
		}
	}
	return val
}

func checkOneChange(s, e int) bool {
	diff := s - e
	if diff < 0 {
		diff = -diff
	}

	var rem, num int
	for diff > 0 {
		rem = diff % 10
		num = diff / 10
		if ((rem != 0 || num <= 0) && (num != 0 || rem <= 0)) || (rem > 3) {
			return false
		}
		diff /= 10
	}
	return true
}
