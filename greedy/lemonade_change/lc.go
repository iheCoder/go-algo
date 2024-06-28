package lemonade_change

func lemonadeChange(bills []int) bool {
	m := make([]int, 3)
	for _, bill := range bills {
		switch bill {
		case 5:
			m[0]++
		case 10:
			if m[0] == 0 {
				return false
			}
			m[0]--
			m[1]++
		case 20:
			if m[0] == 0 || (m[1] == 0 && m[0] < 3) {
				return false
			}

			if m[1] == 0 {
				m[0] -= 3
			} else {
				m[0]--
				m[1]--
			}

			m[2]++
		}
	}

	return true
}
