package buddy_strings

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		seen := [26]bool{}
		for _, ch := range s {
			if seen[ch-'a'] {
				return true
			}
			seen[ch-'a'] = true
		}
		return false
	}

	n := len(s)
	var used int
	var oc, rc byte
	for i := 0; i < n; i++ {
		if s[i] == goal[i] {
			continue
		}

		if used > 1 || used == 1 && (goal[i] != oc || s[i] != rc) {
			return false
		}
		used++
		oc, rc = s[i], goal[i]
	}
	return used == 2
}
