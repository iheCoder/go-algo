package ugly_number

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	uf := []int{2, 3, 5}
	for _, u := range uf {
		for n%u == 0 {
			n /= u
		}
	}
	return n == 1
}
