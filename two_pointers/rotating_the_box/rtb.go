package rotating_the_box

func rotateTheBox(box [][]byte) [][]byte {
	rbRowLen := len(box[0])
	rbColLen := len(box)
	rb := make([][]byte, rbRowLen)
	for j := 0; j < len(rb); j++ {
		rb[j] = make([]byte, rbColLen)
	}

	var i int
	for j, b := range box {
		rbj := rbColLen - j - 1
		for i = len(b) - 1; i >= 0; i-- {
			rb[i][rbj] = b[i]
			if b[i] == '.' {
				k := i - 1
				for ; k >= 0 && b[k] == '.'; k-- {
					rb[k][rbj] = '.'
				}
				if k < 0 {
					break
				}
				if b[k] == '#' {
					b[i] = '#'
					b[k] = '.'
					rb[i][rbj] = '#'
				}
				rb[k][rbj] = b[k]
			}
		}
	}

	return rb
}

// 			if b[i] == '*' {
//				for i >= 0 && (b[i] != '.') {
//					rb[i][j] = b[i]
//					i--
//				}
//				e = i
//				for i = i - 1; i >= 0; i-- {
//					if b[i] == '#' {
//						b[i] = '.'
//						b[e] = '#'
//						rb[e][j] = '#'
//						rb[i][j] = '.'
//						break
//					}
//					rb[i][j] = '.'
//				}
//			} else if b[i] == '.' {
//				e = i
//			} else {
//				b[i] = '.'
//				b[e] = '#'
//				if b[i] == '.' {
//					e = i
//				}
//			}
