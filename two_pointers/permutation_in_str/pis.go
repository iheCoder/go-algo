package permutation_in_str

func checkInclusionTwoMap(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	if len(s1) > len(s2) {
		return false
	}

	var count int
	m1 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
		count++
	}

	m2 := make(map[byte]int)
	var j, k int
	for ; k < len(s2); k++ {
		_, ok := m1[s2[k]]
		if ok {
			m2[s2[k]]++
			if k-j < count-1 {
				continue
			}
			if selfCmpByteMap(m1, m2) {
				return true
			}
			m2[s2[j]]--
			j++
			continue
		}
		m2 = make(map[byte]int)
		j = k + 1
	}

	return false
}

func selfCmpByteMap(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		r, ok := m2[k]
		if !ok {
			return false
		}
		if r != v {
			return false
		}
	}

	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	if len(s1) > len(s2) {
		return false
	}

	//var count int
	//m1 := make(map[byte]int)
	//for i := 0; i < len(s1); i++ {
	//	m1[s1[i]]++
	//	count++
	//}
	//
	//var j, k int
	//m2 := make(map[byte]int)
	//for ; j < len(s2); j++ {
	//
	//}

	return false
}
