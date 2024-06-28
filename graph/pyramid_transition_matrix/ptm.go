package pyramid_transition_matrix

func pyramidTransition(bottom string, allowed []string) bool {
	m := make(map[string][]byte)
	for _, s := range allowed {
		m[s[:2]] = append(m[s[:2]], s[2])
	}

	var dfs func(b string) bool
	dfs = func(b string) bool {
		n := len(b)
		if n == 1 {
			return true
		}

		nb := []string{""}
		tmp := make([]string, 0)
		for i := 1; i < n; i++ {
			vals, ok := m[b[i-1:i+1]]
			if !ok {
				return false
			}

			for j := 0; j < len(nb); j++ {
				for k := 0; k < len(vals); k++ {
					tmp = append(tmp, nb[j]+string(vals[k]))
				}
			}
			nb = tmp
			tmp = []string{}
		}

		for i := 0; i < len(nb); i++ {
			if dfs(nb[i]) {
				return true
			}
		}
		return false
	}
	return dfs(bottom)
}
