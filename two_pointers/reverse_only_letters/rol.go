package reverse_only_letters

func reverseOnlyLetters(s string) string {
	if len(s) <= 1 {
		return s
	}

	rb := make([]byte, len(s))
	i, j := 0, len(s)-1
	for i <= j {
		for !isAlphabet(s[i]) && i < j {
			rb[i] = s[i]
			i++
		}
		for !isAlphabet(s[j]) && i < j {
			rb[j] = s[j]
			j--
		}
		rb[i], rb[j] = s[j], s[i]
		// 用过之后i、j要更新的啊
		i++
		j--
	}

	return string(rb)
}

func isAlphabet(a byte) bool {
	return (65 <= a && a <= 90) || (97 <= a && a <= 122)
}
