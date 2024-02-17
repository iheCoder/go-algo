package expressive_words

func expressiveWords(s string, words []string) int {
	sm := make([]byte, 0)
	var c byte
	var count uint8
	c = s[0]
	count = 1
	for i := 1; i < len(s); i++ {
		if s[i] == c {
			count++
		} else {
			sm = append(sm, c)
			sm = append(sm, count)
			count = 1
			c = s[i]
		}
	}
	sm = append(sm, c)
	sm = append(sm, count)

	var result int
	count = 1
	var k int
	for i := 0; i < len(words); i++ {
		k = 0
		c = words[i][0]
		if sm[k] != c {
			continue
		}
		if len(words[i]) == 1 && len(sm) == 2 && sm[1] == 1 {
			result++
			continue
		}
		count = 1
		for j := 1; j < len(words[i]); j++ {
			if words[i][j] == c {
				count++
			} else {
				c = words[i][j]
				if k+1 > len(sm)-1 {
					break
				}
				// 判断在存在下一个字符时，是否sm的下一个字符与c相同
				if k+2 < len(sm)-1 && c != sm[k+2] {
					break
				}
				// 判断当前字符组是否能扩充为sm所对应的字符组
				if (sm[k+1] < 3 && sm[k+1] != count) || sm[k+1] < count {
					break
				}
				k += 2
				count = 1
			}
			if j == len(words[i])-1 {
				if k+1 == len(sm)-1 && ((sm[k+1] >= 3 && sm[k+1] > count) || sm[k+1] == count) {
					result++
				}
			}
		}
	}

	return result
}
