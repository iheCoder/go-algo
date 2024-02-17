package design_add_and_search_words

import "strings"

type tries struct {
	val      string
	hasEnd   bool
	children []*tries
}

func (t *tries) add(val string) {
	if len(val) == 0 {
		return
	}

	// 提取共同前缀
	cp := extractPrefix(t.val, val)

	// 若共同前缀等于添加值，且等于节点val，则直接设置hasEnd返回
	// 否则新建节点，将公共前缀作为新节点
	if len(cp) == len(val) {
		if len(cp) == len(t.val) {
			t.hasEnd = true
		} else {
			c1 := &tries{
				val:      t.val[len(cp):],
				hasEnd:   t.hasEnd,
				children: t.children,
			}
			t.hasEnd = true
			t.val = cp
			t.children = []*tries{c1}
		}
		return
	}

	// 若共同前缀等于节点val，那么设置添加到子节点中
	if len(cp) == len(t.val) {
		rem := val[len(cp):]
		for _, child := range t.children {
			if child.val[0] == rem[0] {
				child.add(rem)
				return
			}
		}
		t.children = append(t.children, &tries{
			val:    rem,
			hasEnd: true,
		})
		return
	}

	// 若节点val大于共同前缀，则提取新节点val为节点val不同、插入值不同，旧节点设置为共同前缀
	c1 := &tries{
		val:      t.val[len(cp):],
		hasEnd:   t.hasEnd,
		children: t.children,
	}
	c2 := &tries{
		val:    val[len(cp):],
		hasEnd: true,
	}
	t.hasEnd = false
	t.val = cp
	t.children = []*tries{c1, c2}
}

func (t *tries) find(val string) bool {
	if len(val) == 0 {
		return true
	}

	// 直接提取公共前缀，若长度低于树值，则不可能找到
	cp := extractPrefix(t.val, val)
	if len(cp) < len(t.val) {
		return false
	}

	// 如果公共前缀等于目标值，则返回到该树节点是否存在字符串
	if len(cp) == len(val) {
		return t.hasEnd
	}

	// 其他情况，说明仅部分匹配，则取目标值未匹配部分继续匹配
	rem := val[len(cp):]

	// 若剩余首位为.，则尝试匹配所有子节点
	if rem[0] == '.' {
		for _, child := range t.children {
			// 可能存在child val为空的情况吗
			if child.find(rem) {
				return true
			}
		}
		return false
	}

	// 尝试匹配对应子节点
	for _, child := range t.children {
		if child.val[0] == rem[0] {
			return child.find(rem)
		}
	}

	return false
}

func extractPrefix(source, target string) string {
	if len(source) == 0 || len(target) == 0 {
		return ""
	}

	var i int
	for i = 0; i < len(source) && i < len(target); i++ {
		if source[i] != target[i] && target[i] != '.' {
			break
		}
	}
	return source[:i]
}

type WordDictionary struct {
	m map[string]bool
	t *tries
}

func Constructor() WordDictionary {
	return WordDictionary{
		m: make(map[string]bool),
		t: &tries{},
	}
}

func (d *WordDictionary) AddWord(word string) {
	d.m[word] = true
	d.t.add(word)
}

func (d *WordDictionary) Search(word string) bool {
	if d.m[word] {
		return true
	}

	if !strings.Contains(word, ".") {
		return false
	}

	return d.t.find(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// 	if len(t.val) > 0 {
//		// 还存在一种情况，如果是在abc搜索a.，那么前缀也会等于2
//		if len(cp) == 0 {
//			return false
//		}
//		if len(cp) == len(val) {
//			return false
//			//return len(cp) == len(t.val) && t.hasEnd
//		}
//		rem = val[len(cp):]
//	}

// func checkEqual(source, target string) bool {
//	if len(source) != len(target) {
//		return false
//	}
//
//	if !strings.Contains(target, ".") {
//		return source == target
//	}
//
//	var i int
//	for i = 0; i < len(source); i++ {
//		if source[i] != target[i] && target[i] != '.' {
//			return false
//		}
//	}
//	return true
//}
