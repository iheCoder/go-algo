package bitwise_and_of_numbers_range

func rangeBitwiseAnd(left int, right int) int {
	if left == 0 {
		return 0
	}
	if left == right {
		return right
	}

	xor := left ^ right
	hb := getHighest1Bit(xor)
	return right &^ (1<<(hb+1) - 1)
}

func getHighest1Bit(x int) int {
	var pos int
	if x&0xffff0000 > 0 {
		pos += 16
		x >>= 16
	}
	if x&0xff00 > 0 {
		pos += 8
		x >>= 8
	}
	if x&0xf0 > 0 {
		pos += 4
		x >>= 4
	}
	if x&0xc > 0 {
		pos += 2
		x >>= 2
	}
	if x&0x2 > 0 {
		pos += 1
	}
	return pos
}
