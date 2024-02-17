package bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	sm := make([]int, 10)
	gm := make([]int, 10)

	var ac, bc int
	n := len(secret)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			ac++
		}

		sm[convertByteToInt(secret[i])]++
		gm[convertByteToInt(guess[i])]++
	}

	for i := 0; i < 10; i++ {
		bc += minInt(sm[i], gm[i])
	}
	bc -= ac

	return fmt.Sprintf("%dA%dB", ac, bc)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func convertByteToInt(b byte) int {
	switch b {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return 0
}
