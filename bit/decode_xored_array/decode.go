package decode_xored_array

func decode(encoded []int, first int) []int {
	l := len(encoded) + 1
	r := make([]int, l)
	r[0] = first
	for i := 1; i < l; i++ {
		r[i] = r[i-1] ^ encoded[i-1]
	}
	return r
}
