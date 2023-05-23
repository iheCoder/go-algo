package reverse_vowels_of_a_string

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}

	vm := map[byte]bool{
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

	bs := []byte(s)
	i, j := 0, len(s)-1
	for {
		for i < j && !vm[bs[i]] {
			i++
		}
		if i >= j {
			return string(bs)
		}
		for i < j && !vm[bs[j]] {
			j--
		}
		if i >= j {
			return string(bs)
		}
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}
