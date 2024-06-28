package swap_adjacent_in_lr_string

var transitions = map[string]string{
	"XL": "LX",
	"RX": "XR",
}

func canTransform(start string, end string) bool {
	n := len(start)
	var i, j int
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}

		if i < n && j < n {
			if start[i] != end[j] {
				return false
			}

			if start[i] == 'L' && i < j || start[i] == 'R' && i > j {
				return false
			}

			i++
			j++
		}
	}

	for i < n {
		if start[i] != 'X' {
			return false
		}
		i++
	}

	for j < n {
		if end[j] != 'X' {
			return false
		}
		j++
	}

	return true
}

// 	if n == 1 {
//		return start == end
//	}
//
//	i := 1
//	sb := make([]byte, n)
//	copy(sb, start)
//	for i < n {
//		if sb[i-1] != end[i-1] {
//			r, ok := transitions[string(sb[i-1:i+1])]
//			if !ok || r != end[i-1:i+1] {
//				return false
//			}
//			sb[i-1] = r[0]
//			sb[i] = r[1]
//		}
//		i++
//	}
//	return sb[n-1] == end[n-1]
