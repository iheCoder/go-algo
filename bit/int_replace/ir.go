package int_replace

func integerReplacement(n int) int {
	var c int
	for n > 1 {
		if n&1 == 0 {
			n >>= 1
		} else {
			if n != 3 && (n+1)&-(n+1) > (n-1)&-(n-1) {
				n++
			} else {
				n--
			}

		}
		c++
	}
	return c
}

//if n != 3 && n&(n+1) == 0 {
//	n++
//} else {
//	n--
//}
