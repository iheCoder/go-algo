package reconstruct_original_digits_from_english

import "strings"

var wdm = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func originalDigits(s string) string {
	sb := make([][]byte, 10)
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for m['z'] > 0 && m['e'] > 0 && m['r'] > 0 && m['o'] > 0 {
		sb[0] = append(sb[0], '0')
		m['z']--
		m['e']--
		m['r']--
		m['o']--
	}

	for m['s'] > 0 && m['i'] > 0 && m['x'] > 0 {
		sb[6] = append(sb[6], '6')
		m['s']--
		m['i']--
		m['x']--
	}

	for m['s'] > 0 && m['e'] > 1 && m['v'] > 0 && m['n'] > 0 {
		sb[7] = append(sb[7], '7')
		m['s']--
		m['e'] -= 2
		m['v']--
		m['n']--
	}

	for m['e'] > 0 && m['i'] > 0 && m['g'] > 0 && m['h'] > 0 && m['t'] > 0 {
		sb[8] = append(sb[8], '8')
		m['e']--
		m['i']--
		m['g']--
		m['h']--
		m['t']--
	}

	for m['t'] > 0 && m['e'] > 1 && m['h'] > 0 && m['r'] > 0 {
		sb[3] = append(sb[3], '3')
		m['t']--
		m['e'] -= 2
		m['r']--
		m['h']--
	}

	for m['f'] > 0 && m['u'] > 0 && m['r'] > 0 && m['o'] > 0 {
		sb[4] = append(sb[4], '4')
		m['f']--
		m['u']--
		m['r']--
		m['o']--
	}

	for m['f'] > 0 && m['e'] > 0 && m['i'] > 0 && m['v'] > 0 {
		sb[5] = append(sb[5], '5')
		m['f']--
		m['i']--
		m['v']--
		m['e']--
	}

	for m['n'] > 1 && m['e'] > 0 && m['i'] > 0 {
		sb[9] = append(sb[9], '9')
		m['n'] -= 2
		m['i']--
		m['e']--
	}

	for m['e'] > 0 && m['n'] > 0 && m['o'] > 0 {
		sb[1] = append(sb[1], '1')
		m['e']--
		m['n']--
		m['o']--
	}

	for m['t'] > 0 && m['w'] > 0 && m['o'] > 0 {
		sb[2] = append(sb[2], '2')
		m['t']--
		m['w']--
		m['o']--
	}

	r := &strings.Builder{}
	for _, bytes := range sb {
		r.Write(bytes)
	}
	return r.String()
}
