package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	var i, j, k int
	r := make([]byte, len(word1)+len(word2))
	for {
		if j >= len(word1) {
			for k < len(word2) {
				r[i] = word2[k]
				k++
				i++
			}
			return string(r)
		}
		if k >= len(word2) {
			for j < len(word1) {
				r[i] = word1[j]
				j++
				i++
			}
			return string(r)
		}

		r[i] = word1[j]
		i++
		j++
		r[i] = word2[k]
		i++
		k++
	}
}
