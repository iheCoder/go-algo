package push_dominoes

func pushDominoes(dominoes string) string {
	if len(dominoes) <= 1 {
		return dominoes
	}

	ds := []byte(dominoes)
	forces := make([][]int, len(ds))
	for i := 0; i < len(forces); i++ {
		forces[i] = make([]int, 2)
	}

	q := make([]int, 0)
	for i := 0; i < len(ds); i++ {
		if ds[i] != '.' {
			q = append(q, i)
			if i > 0 && ds[i] == 'L' {
				forces[i-1][0]++
			} else if i < len(ds)-1 && ds[i] == 'R' {
				forces[i+1][0]++
			}
		}
	}

	var i int
	termCount := len(q)
	prevForceIndex := 0
	diffForceIndex := 1
	for len(q) != 0 {
		if termCount <= 0 {
			termCount = len(q)
			for j := 0; j < len(forces); j++ {
				forces[j][prevForceIndex] += forces[j][diffForceIndex]
				forces[j][diffForceIndex] = 0
			}
		}
		i = q[0]
		q = q[1:]

		termCount--

		if forces[i][prevForceIndex] > 1 {
			continue
		}

		if ds[i] == 'L' && i > 0 && ds[i-1] == '.' && forces[i-1][prevForceIndex] < 2 {
			if i-2 >= 0 {
				forces[i-2][diffForceIndex]++
			}
			ds[i-1] = 'L'
			q = append(q, i-1)
		} else if ds[i] == 'R' && i < len(ds)-1 && ds[i+1] == '.' && forces[i+1][prevForceIndex] < 2 {
			if i+2 < len(ds) {
				forces[i+2][diffForceIndex]++
			}
			ds[i+1] = 'R'
			q = append(q, i+1)
		}
	}

	return string(ds)
}

// func pushDominoes(dominoes string) string {
//	if len(dominoes) <= 1 {
//		return dominoes
//	}
//
//	ds := []byte(dominoes)
//	forces := make([][]byte, len(ds))
//
//	q := make([]int, 0)
//	for i := 0; i < len(ds); i++ {
//		if ds[i] != '.' {
//			q = append(q, i)
//			if i > 0 && ds[i] == 'L' {
//				forces[i-1] = append(forces[i-1], ds[i])
//			} else if i < len(ds)-1 && ds[i] == 'R' {
//				forces[i+1] = append(forces[i+1], ds[i])
//			}
//			//forces[i][1] = ds[i]
//			//forces[i+2][0] = ds[i]
//		}
//	}
//
//	var i int
//	for len(q) != 0 {
//		i = q[0]
//		q = q[1:]
//
//		if len(forces[i]) > 1 {
//			continue
//		}
//
//		// 问题在于同一秒内的操作应该是并行的，可是如果在每一次都改变了受力这就会影响同秒内的下次操作
//		if ds[i] == 'L' && i > 0 && ds[i-1] == '.' && len(forces[i-1]) < 2 {
//			if i-2 >= 0 {
//				forces[i-2] = append(forces[i-2], 'L')
//			}
//			ds[i-1] = 'L'
//			q = append(q, i-1)
//			//forces[i][1] = 'L'
//		} else if ds[i] == 'R' && i < len(ds)-1 && ds[i+1] == '.' && len(forces[i+1]) < 2 {
//			if i+2 < len(ds) {
//				forces[i+2] = append(forces[i+2], 'R')
//			}
//			ds[i+1] = 'R'
//			q = append(q, i+1)
//			//forces[i+2][0] = 'R'
//		}
//	}
//
//	return string(ds)
//}
