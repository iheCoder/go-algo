package xor_op_in_an_array

func xorOperation(n int, start int) int {
	if n == 1 {
		return start
	}

	var r int
	for i := 0; i < n; i++ {
		r ^= start + i<<1
	}
	return r
}
