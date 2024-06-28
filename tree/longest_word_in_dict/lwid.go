package longest_word_in_dict

type trie struct {
	isEnd    bool
	children []*trie
}

func (t *trie) add(val string) {
	if len(val) == 0 {
		t.isEnd = true
		return
	}

	i := int(val[0] - 'a')
	if t.children[i] == nil {
		t.children[i] = &trie{children: make([]*trie, 26)}
	}

	t.children[i].add(val[1:])
}

func (t *trie) find() string {
	var ms string
	for i, child := range t.children {
		if child == nil {
			continue
		}
		if !child.isEnd {
			continue
		}
		cf := string(rune(i+'a')) + child.find()
		if len(cf) > len(ms) {
			ms = cf
		}
	}
	return ms
}

func longestWord(words []string) string {
	t := &trie{
		isEnd:    true,
		children: make([]*trie, 26),
	}
	for _, word := range words {
		t.add(word)
	}
	return t.find()
}
