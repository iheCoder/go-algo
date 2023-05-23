package reverse_words_in_a_string_iii

func reverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}

	bs := []byte(s)

	var i, j int
	for start := 0; start < len(s); start++ {
		if s[start] != ' ' {
			i = start
			for start < len(s) && s[start] != ' ' {
				start++
			}
			j = start - 1
			for i < j {
				bs[i], bs[j] = bs[j], bs[i]
				i++
				j--
			}
		}
	}

	return string(bs)
}
