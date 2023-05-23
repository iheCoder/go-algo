package reverse_prefix_of_word

func reversePrefix(word string, ch byte) string {
	var j int
	for ; j < len(word); j++ {
		if word[j] == ch {
			break
		}
	}
	if j >= len(word) {
		return word
	}

	wb := []byte(word)
	var i int
	for i < j {
		wb[i], wb[j] = wb[j], wb[i]
		i++
		j--
	}

	return string(wb)
}
