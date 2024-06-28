package number_of_matching_subsequences

func numMatchingSubseq(s string, words []string) int {
	var ans int
	for _, word := range words {
		var j int
		match := true
		for i := 0; i < len(word); i++ {
			for j < len(s) && s[j] != word[i] {
				j++
			}
			if j >= len(s) {
				match = false
				break
			}
			j++
		}
		if match {
			ans++
		}
	}
	return ans
}

func numMatchingSubseqMulPointer(s string, words []string) int {
	var ans int
	type pair struct {
		// i代表words的索引，j代表在words[i]中的位置
		// j并且也代表与s完全匹配的数量
		i, j int
	}
	ps := [26][]*pair{}
	for i, word := range words {
		ps[word[0]-'a'] = append(ps[word[0]-'a'], &pair{i, 0})
	}

	for _, c := range s {
		// 取出s当前元素的匹配的words
		q := ps[c-'a']
		ps[c-'a'] = nil
		for _, p := range q {
			// 能取出来，说明至少当前匹配了，所以p.j++
			p.j++
			// 当匹配数量与words[i]长度一致，则说明匹配成功
			if p.j == len(words[p.i]) {
				ans++
			} else {
				// 否则就添加待匹配元素到ps中
				w := words[p.i][p.j] - 'a'
				ps[w] = append(ps[w], p)
			}
		}
	}

	return ans
}
