package bulb_switcher_ii

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}

	if n == 1 {
		return 2
	} else if n == 2 {
		if presses == 1 {
			return 3
		}
		return 4
	}

	switch presses {
	case 1:
		return 4
	case 2:
		return 7
	default:
		return 8
	}
}
