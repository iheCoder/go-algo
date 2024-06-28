package replace_words

import "strings"

type tries struct {
	val      string
	children []*tries
	isEnd    bool
}

func (t *tries) add(val string) {
	// 1. 若插入val恰好为为节点val，则设置当前节点可到达
	if val == t.val {
		t.isEnd = true
		return
	}

	// 2. 取共同前缀
	cp := getCommonPrefix(t.val, val)

	// 3. 若共同前缀恰好为val，则中间插入新节点
	if cp == val {
		nt := &tries{
			val:      t.val[len(cp):],
			children: t.children,
			isEnd:    true,
		}
		t.val = t.val[:len(cp)]
		t.children = []*tries{nt}
		return
	}

	// 3. 若共同前缀恰好为当前节点val，那么直接插入子节点
	if cp == t.val || len(t.val) == 0 {
		val = val[len(cp):]
		for _, child := range t.children {
			if child.val[0] == val[0] {
				child.add(val)
				return
			}
		}
		t.children = append(t.children, &tries{
			val:   val,
			isEnd: true,
		})
		return
	}

	// 4. 若都不是，则中间插入新节点
	nt1 := &tries{
		val:   val[len(cp):],
		isEnd: true,
	}
	nt2 := &tries{
		val:      t.val[len(cp):],
		children: t.children,
		isEnd:    t.isEnd,
	}
	t.val = cp
	t.children = []*tries{nt1, nt2}
	t.isEnd = false
}

func (t *tries) search(val string) string {
	return t.find(val, "")
}

func (t *tries) find(val, cur string) string {
	if t.isEnd {
		return cur
	}

	for _, child := range t.children {
		cp := getCommonPrefix(child.val, val)
		if len(cp) == len(child.val) {
			return child.find(val[len(cp):], cur+cp)
		}
	}
	return ""
}

func getCommonPrefix(s1, s2 string) string {
	var i, j int
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			break
		}
		i++
		j++
	}

	return s1[:i]
}

func replaceWords(dictionary []string, sentence string) string {
	t := &tries{}
	for _, s := range dictionary {
		t.add(s)
	}

	words := strings.Split(sentence, " ")
	ans := make([]string, len(words))
	for i, word := range words {
		rw := t.search(word)
		if len(rw) > 0 {
			ans[i] = rw
		} else {
			ans[i] = word
		}
	}

	return strings.Join(ans, " ")
}
