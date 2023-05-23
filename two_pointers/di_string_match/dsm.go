package di_string_match

func diStringMatch(s string) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}

	r := make([]int, n+1)
	low, high := 0, n
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			r[i] = low
			low++
		} else {
			r[i] = high
			high--
		}
	}
	r[n] = low
	return r
}

// 简单题。这么多代码还是错的
func diStringMatchDep(s string) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}

	r := make([]int, n+1)
	m := make(map[int]bool)
	for i := 0; i < n+1; i++ {
		m[i] = false
	}

	start := (n + 1) / 2
	m[start] = true
	r[0] = start
	var cur int
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			cur = r[i] + 1
			for m[cur] {
				cur++
			}
			if cur > n {
				cur = n
				for j := i; j >= 0; j-- {
					m[r[j]] = false
					r[j]--
				}
				for j := 0; j <= i+1; j++ {
					m[r[j]] = true
				}
			}
		} else {
			cur = r[i] - 1
			for m[cur] {
				cur--
			}
			if cur < 0 {
				cur = 0
				for j := i; j >= 0; j-- {
					m[r[j]] = false
					r[j]++
				}
				for j := 0; j <= i+1; j++ {
					m[r[j]] = true
				}
			}
		}
		r[i+1] = cur
		m[cur] = true
	}

	return r
}

// 	var count int
//	for i := 0; i < n; i++ {
//		if s[i] == 'I' {
//			r[i+1] = r[i] + 1
//		} else {
//			r[i+1] = r[i] - 1
//			count++
//		}
//	}
//
//	for i := 0; i < len(r); i++ {
//		r[i]+=count
//	}
