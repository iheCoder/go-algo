package map_sum_pairs

type trie struct {
	children []*trie
	isEnd    bool
	sum      int
}

func (t *trie) add(key string, val int) {
	if len(key) == 0 {
		t.isEnd = true
		t.sum += val
		return
	}

	i := key[0] - 'a'
	if t.children[i] == nil {
		child := &trie{
			children: make([]*trie, 26),
			isEnd:    true,
		}
		t.children[i] = child
	}

	t.sum += val
	t.children[i].add(key[1:], val)
}

func (t *trie) search(prefix string) int {
	if len(prefix) == 0 {
		return t.sum
	}

	i := int(prefix[0] - 'a')
	if t.children[i] == nil {
		return 0
	}

	return t.children[i].search(prefix[1:])
}

type MapSum struct {
	t *trie
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{
		t: &trie{children: make([]*trie, 26)},
		m: make(map[string]int),
	}
}

func (s *MapSum) Insert(key string, val int) {
	diff := val
	if v, ok := s.m[key]; ok {
		diff = val - v
	}

	s.m[key] = val
	s.t.add(key, diff)
}

func (s *MapSum) Sum(prefix string) int {
	return s.t.search(prefix)
}
