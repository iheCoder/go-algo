package push_dominoes

func pushDominoesDep(dominoes string) string {
	if len(dominoes) <= 1 {
		return dominoes
	}

	ds := []byte(dominoes)
	stack := make([]int, 0)
	for i := 0; i < len(ds); i++ {
		if ds[i] == '.' {
			stack = append(stack, i)
			for ; i < len(ds) && ds[i] == '.'; i++ {
			}
			if i >= len(ds) && ds[len(ds)-1] == '.' {
				stack = append(stack, len(ds)-1)
			} else {
				stack = append(stack, i-1)
			}
		}
	}
	if len(stack) == 0 {
		return dominoes
	}

	var ls1, rs1 byte
	var ls2, rs2 byte
	var i int
	n := len(ds)
	var j, k int
	for i+1 < len(stack) {
		j, k = stack[i], stack[i+1]
		for j <= k {
			ls1, rs1 = getStatus(j-1, n, ds), getStatus(j+1, n, ds)
			ls2, rs2 = getStatus(k-1, n, ds), getStatus(k+1, n, ds)
			ds[j] = alterStatus(ls1, rs1)
			ds[k] = alterStatus(ls2, rs2)
			j++
			k--
		}
		i += 2
	}
	if ds[0] == '.' {
		for i = 1; i < len(ds) && ds[i] == '.'; i++ {
		}
		if i < len(ds) && ds[i] == 'L' {
			for i = i - 1; i >= 0; i-- {
				ds[i] = 'L'
			}
		}
	}
	if ds[len(ds)-1] == '.' {
		for i = len(ds) - 2; i >= 0 && ds[i] == '.'; i-- {
		}
		if i >= 0 && ds[i] == 'R' {
			for i = i + 1; i < len(ds); i++ {
				ds[i] = 'R'
			}
		}
	}

	return string(ds)
}

func getStatus(index, n int, bs []byte) byte {
	if index < 0 || index >= n {
		return '.'
	}
	return bs[index]
}

func alterStatus(ls, rs byte) byte {
	if ls != 'R' && rs == 'L' {
		return 'L'
	} else if ls == 'R' && rs != 'L' {
		return 'R'
	}

	return '.'
}
