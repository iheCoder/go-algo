package longest_word_in_dictionary_through_deleting

import "sort"

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	var r string
	var j, k int
	for i := len(dictionary) - 1; i >= 0; i-- {
		if len(dictionary[i]) > len(s) {
			continue
		} else if len(dictionary[i]) == len(s) {
			if dictionary[i] == s {
				return s
			}
			continue
		}

		j, k = 0, 0
		for j < len(dictionary[i]) {
			for ; k < len(s); k++ {
				if dictionary[i][j] == s[k] {
					k++
					j++
					break
				}
			}
			if k >= len(s) {
				break
			}
		}
		if j == len(dictionary[i]) {
			r = minStr(r, dictionary[i])
		}

		if i == 0 || len(dictionary[i-1]) < len(r) {
			return r
		}
	}

	return r
}

func minStr(s1, s2 string) string {
	if len(s1) == 0 {
		return s2
	}

	if s1 < s2 {
		return s1
	}
	return s2
}

// 				// 会不会存在虽然j++，但是由于已经超出k的长度，所以就break，但是此时j中还没有对应匹配的
//				// 不可能。若能进入该循环，说明j是小于字符串dictionary长度的。所以j
// 其实应该是最后一个dict元素没有在s中找到，那么此时j还是会++，于是还是会等于len(dict[i])
