package add_binary

func addBinary(a string, b string) string {
	var r []byte
	ai, bi := len(a)-1, len(b)-1
	e := 0
	var sum int
	for ai >= 0 || bi >= 0 {
		if ai >= 0 && a[ai] == '1' {
			sum++
		}
		if bi >= 0 && b[bi] == '1' {
			sum++
		}
		if e == 1 {
			sum++
		}
		e = 0
		if sum > 1 {
			e = 1
		}
		if sum == 0 || sum == 2 {
			r = append(r, '0')
		} else {
			r = append(r, '1')
		}
		sum = 0
		ai--
		bi--
	}
	if e == 1 {
		r = append(r, '1')
	}

	i, j := 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}
