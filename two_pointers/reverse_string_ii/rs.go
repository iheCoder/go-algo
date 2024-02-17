package reverse_string_ii

func reverseStr(s string, k int) string {
	bs := []byte(s)
	if k >= len(bs) {
		reverseBytes(bs)
		return string(bs)
	}

	start, end := 0, k
	for {
		reverseBytes(bs[start:end])
		start = start + 2*k
		if start >= len(bs) {
			break
		}
		end = start + k
		if end > len(bs) {
			end = len(bs)
		}
	}

	return string(bs)
}

func reverseBytes(bs []byte) {
	if len(bs) <= 1 {
		return
	}

	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}
