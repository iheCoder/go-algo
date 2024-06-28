package shifting_letters

func shiftingLetters(s string, shifts []int) string {
	n := len(s)
	sums := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sums[i] = (sums[i+1] + shifts[i]) % 26
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = movLetter(s[i], byte(sums[i]))
	}

	return string(res)
}

func movLetter(b, mv byte) byte {
	r := b + mv
	if r <= 'z' {
		return r
	}
	return r - 'z' - 1 + 'a'
}
