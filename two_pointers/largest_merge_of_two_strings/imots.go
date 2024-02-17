package largest_merge_of_two_strings

func largestMerge(word1 string, word2 string) string {
	r := make([]byte, len(word1)+len(word2))
	var i, j, k int
	for ; ; i++ {
		if j >= len(word1) {
			for k < len(word2) {
				r[i] = word2[k]
				i++
				k++
			}
			break
		}
		if k >= len(word2) {
			for j < len(word1) {
				r[i] = word1[j]
				i++
				j++
			}
			break
		}

		if word1[j] > word2[k] {
			r[i] = word1[j]
			j++
		} else if word1[j] < word2[k] {
			r[i] = word2[k]
			k++
		} else {
			// 仅仅往后看一位是不对的！！！
			r[i] = word1[j]
			i++
			if j+1 < len(word1) && k+1 < len(word2) {
				if word1[j] >= word1[j+1] && word1[j] >= word2[k+1] {
					j++
					r[i] = word2[k]
					k++
				} else if word1[j+1] >= word1[j] && word1[j+1] >= word2[k+1] {
					j++
					r[i] = word1[j]
					j++
				} else {
					k++
					r[i] = word2[k]
					k++
				}
			} else if j+1 >= len(word1) && i+1 >= len(word2) {
				j++
				r[i] = word2[k]
				k++
				break
			} else if j+1 >= len(word1) {
				if word1[j] >= word2[k+1] {
					j++
					r[i] = word1[k]
					k++
				} else {
					k++
					r[i] = word2[k]
					k++
				}
			} else {
				if word1[k] >= word2[j+1] {
					k++
					r[i] = word1[j]
					j++
				} else {
					j++
					r[i] = word2[j]
					j++
				}
			}
		}
	}

	return string(r)
}

// 					if word1[j] >= word2[k] {
//						r[i] = word1[j]
//						j++
//					} else {
//						r[i] = word1[k]
//						k++
//					}
