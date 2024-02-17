package guess_number_higher_or_lower

func guessNumber(n int) int {
	l, h := 1, n
	x := l + (h-l)/2
	for {
		switch guess(x) {
		case -1:
			h = x - 1
		case 1:
			l = x + 1
		default:
			return x
		}
		x = l + (h-l)/2
	}
}

func guess(num int) int {
	return 0
}
