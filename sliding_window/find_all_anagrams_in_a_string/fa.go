package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	pm := make(map[byte]int)
	n := len(p)
	for i := 0; i < n; i++ {
		pm[p[i]]++
	}
	var r []int
	sm := make(map[byte]int)
	for i := 0; i < n; i++ {
		sm[s[i]]++
	}
	if isEqual(pm, sm) {
		r = append(r, 0)
	}

	i, j := 0, n
	for j < len(s) {
		sm[s[i]]--
		if sm[s[i]] == 0 {
			delete(sm, s[i])
		}
		sm[s[j]]++
		i++
		j++

		if isEqual(pm, sm) {
			r = append(r, i)
		}
	}
	return r
}

func isEqual(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for b, i := range m1 {
		if i1, ok := m2[b]; !ok {
			return false
		} else if i1 != i {
			return false
		}
	}
	return true
}
