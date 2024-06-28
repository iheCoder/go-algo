package implement_magic_dict

type trie struct {
	children []*trie
	isEnd    bool
}

func (t *trie) add(val string) {
	if len(val) == 0 {
		t.isEnd = true
		return
	}

	i := val[0] - 'a'
	if t.children[i] == nil {
		child := &trie{children: make([]*trie, 26), isEnd: true}
		t.children[i] = child
	}

	t.children[i].add(val[1:])
}

func (t *trie) search(val string, used bool) bool {
	if len(val) == 0 {
		return used && t.isEnd
	}

	i := int(val[0] - 'a')
	if t.children[i] != nil && t.children[i].search(val[1:], used) {
		return true
	}

	if used {
		return false
	}

	used = true
	for j, child := range t.children {
		if child == nil || i == j {
			continue
		}
		if child.search(val[1:], used) {
			return true
		}
	}
	return false
}

type MagicDictionary struct {
	m map[int]*trie
}

func Constructor() MagicDictionary {
	return MagicDictionary{m: make(map[int]*trie)}
}

func (d *MagicDictionary) BuildDict(dictionary []string) {
	for _, dict := range dictionary {
		dl := len(dict)
		if d.m[dl] == nil {
			d.m[dl] = &trie{children: make([]*trie, 26)}
		}
		d.m[dl].add(dict)
	}
}

func (d *MagicDictionary) Search(searchWord string) bool {
	sl := len(searchWord)
	if d.m[sl] == nil {
		return false
	}

	return d.m[sl].search(searchWord, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
