package goat_latin

import "strings"

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func toGoatLatin(sentence string) string {
	ss := strings.Split(sentence, " ")
	n := len(ss)
	ans := make([]string, 0, n)
	for i := 0; i < n; i++ {
		x := []byte(ss[i])
		if !vowels[x[0]] {
			y := x[0]
			x = x[1:]
			x = append(x, y)
		}

		x = append(x, 'm')
		x = append(x, 'a')

		for j := 0; j <= i; j++ {
			x = append(x, 'a')
		}
		ans = append(ans, string(x))
	}

	return strings.Join(ans, " ")
}
