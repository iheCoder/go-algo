package vowel_spellchecker

import "strings"

// 关键是vowel，大小写还好，完全可以用hash搞定

func spellchecker(wordlist []string, queries []string) []string {
	wordsPerfect := make(map[string]bool)
	wordsCap := make(map[string]string)
	wordsVow := make(map[string]string)

	for _, word := range wordlist {
		wordsPerfect[word] = true

		lw := strings.ToLower(word)
		if _, exists := wordsCap[lw]; !exists {
			wordsCap[lw] = word
		}

		vw := devowel(lw)
		if _, exists := wordsVow[vw]; !exists {
			wordsVow[vw] = word
		}
	}

	solve := func(q string) string {
		if wordsPerfect[q] {
			return q
		}

		lq := strings.ToLower(q)
		if x, exists := wordsCap[lq]; exists {
			return x
		}

		vq := devowel(lq)
		if x, exists := wordsVow[vq]; exists {
			return x
		}

		return ""
	}

	ans := make([]string, 0, len(queries))
	for _, query := range queries {
		ans = append(ans, solve(query))
	}

	return ans
}

func devowel(word string) string {
	var sb strings.Builder
	for _, c := range word {
		if isVowel(c) {
			sb.WriteRune('*')
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

//func spellchecker(wordlist []string, queries []string) []string {
//	fm := make(map[string]bool)
//	lcm := make(map[string]string)
//	vm := make(map[string]string)
//
//	for _, s := range wordlist {
//		fm[s] = true
//		ls := strings.ToLower(s)
//		lcm[ls] = s
//
//		a := genVowelPossible(ls)
//		for _, x := range a {
//			vm[x] = s
//		}
//	}
//
//	ans := make([]string, 0, len(queries))
//	for _, query := range queries {
//		if fm[query] {
//			ans = append(ans, query)
//			continue
//		}
//
//		lq := strings.ToLower(query)
//		x, ok := lcm[lq]
//		if ok {
//			ans = append(ans, x)
//			continue
//		}
//
//		y, ok := vm[lq]
//		if ok {
//			ans = append(ans, y)
//			continue
//		}
//
//		ans = append(ans, "")
//	}
//
//	return ans
//}

func genVowelPossible(s string) []string {
	var ans [][]byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			tmp := make([][]byte, 0, 5*len(ans))
			for j := 0; j < len(ans); j++ {
				tmp = append(tmp, append(ans[j], 'a'))
				tmp = append(tmp, append(ans[j], 'e'))
				tmp = append(tmp, append(ans[j], 'i'))
				tmp = append(tmp, append(ans[j], 'o'))
				tmp = append(tmp, append(ans[j], 'u'))
			}

		default:
			for j, _ := range ans {
				ans[j] = append(ans[j], s[i])
			}
		}
	}

	r := make([]string, len(ans))
	for i, an := range ans {
		r[i] = string(an)
	}

	return r
}
