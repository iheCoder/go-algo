package short_encoding_of_words

//type tries struct {
//	val      string
//	hasEnd   bool
//	children []*tries
//}
//
//func (t *tries) add(val string) {
//	if len(val) == 0 {
//		return
//	}
//
//	// 提取共同前缀
//	cp := extractPrefix(t.val, val)
//
//	// 若共同前缀等于添加值，且等于节点val，则直接设置hasEnd返回
//	// 否则新建节点，将公共前缀作为新节点
//	if len(cp) == len(val) {
//		if len(cp) == len(t.val) {
//			t.hasEnd = true
//		} else {
//			c1 := &tries{
//				val:      t.val[len(cp):],
//				hasEnd:   t.hasEnd,
//				children: t.children,
//			}
//			t.hasEnd = true
//			t.val = cp
//			t.children = []*tries{c1}
//		}
//		return
//	}
//
//	// 若共同前缀等于节点val，那么设置添加到子节点中
//	if len(cp) == len(t.val) {
//		rem := val[len(cp):]
//		for _, child := range t.children {
//			if child.val[0] == rem[0] {
//				child.add(rem)
//				return
//			}
//		}
//		t.children = append(t.children, &tries{
//			val:    rem,
//			hasEnd: true,
//		})
//		return
//	}
//
//	// 若节点val大于共同前缀，则提取新节点val为节点val不同、插入值不同，旧节点设置为共同前缀
//	c1 := &tries{
//		val:      t.val[len(cp):],
//		hasEnd:   t.hasEnd,
//		children: t.children,
//	}
//	c2 := &tries{
//		val:    val[len(cp):],
//		hasEnd: true,
//	}
//	t.hasEnd = false
//	t.val = cp
//	t.children = []*tries{c1, c2}
//}
//
//func extractPrefix(source, target string) string {
//	if len(source) == 0 || len(target) == 0 {
//		return ""
//	}
//
//	var i int
//	for i = 0; i < len(source) && i < len(target); i++ {
//		if source[i] != target[i] {
//			break
//		}
//	}
//	return source[:i]
//}
//
//func (t *tries) countOfPath(vp int) int {
//	vp += len(t.val)
//	if len(t.children) == 0 {
//		return vp + 1
//	}
//
//	var total int
//	for _, child := range t.children {
//		total += child.countOfPath(vp)
//	}
//	return total
//}

func minimumLengthEncoding(words []string) int {
	m := make(map[string]bool)
	for _, word := range words {
		m[word] = true
	}

	for _, word := range words {
		for i := 1; i < len(word); i++ {
			if m[word[i:]] {
				delete(m, word[i:])
			}
		}
	}

	var ans int
	for s, _ := range m {
		ans += len(s) + 1
	}
	return ans
}

// 压缩后缀，可不能压缩任意子字符串啊！
// 	m := make([]string, 1, len(words))
//	m[0] = words[0]
//	for i := 1; i < len(words); i++ {
//		var replaced bool
//		for j := 0; j < len(m); j++ {
//			if len(m[j]) >= len(words[i]) {
//				if strings.Contains(m[j], words[i]) {
//					replaced = true
//					break
//				}
//			} else {
//				if strings.Contains(words[i], m[j]) {
//					m[j] = words[i]
//					replaced = true
//					break
//				}
//			}
//		}
//		if !replaced {
//			m = append(m, words[i])
//		}
//	}
