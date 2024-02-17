package binary_num_with_alternating_bits

const m1 = 0x55555555

func hasAlternatingBits(n int) bool {
	if n == 1 {
		return true
	}
	x := n>>1 ^ n
	return x&(x+1) == 0
}
