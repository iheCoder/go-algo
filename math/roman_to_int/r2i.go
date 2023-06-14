package roman_to_int

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var r, x int
	for i := 0; i < len(s); i++ {
		x = m[s[i]]
		r += x
		if i > 0 && x > m[s[i-1]] {
			r -= 2 * m[s[i-1]]
		}
	}
	return r
}
