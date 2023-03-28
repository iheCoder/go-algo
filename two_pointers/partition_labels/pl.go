package partition_labels

// 先二分为只出现一次的字符串，然后接着在分，直到不能分为止。如果用map内存消耗太大了

func partitionLabels(s string) []int {
	if len(s) == 1 {
		return []int{1}
	}

	m := make(map[byte]int)
	is := make([]int, 0)
	for i := 0; i < len(s); i++ {
		j, ok := m[s[i]]
		if ok {
			k := len(is) - 1
			for ; k >= 0; k-- {
				if is[k] <= j {
					break
				}
			}
			is = is[:k+1]
			continue
		}
		m[s[i]] = i
		is = append(is, i)
	}
	is = append(is, len(s))

	r := make([]int, len(is)-1)
	for i := 1; i < len(is); i++ {
		r[i-1] = is[i] - is[i-1]
	}
	return r
}
