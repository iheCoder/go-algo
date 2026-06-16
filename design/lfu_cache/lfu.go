package lfu_cache

type LFUCache struct {
	capacity int
	size     int
	minFreq  int

	nodes map[int]*Node
	lists map[int]*DList
}

type Node struct {
	key, val   int
	freq       int
	prev, next *Node
}

type DList struct {
	head, tail *Node
	size       int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes:    make(map[int]*Node),
		lists:    make(map[int]*DList),
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.nodes[key]
	if !ok {
		return -1
	}

	c.increaseFreq(node)
	return node.val
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}

	if node, ok := c.nodes[key]; ok {
		node.val = value
		c.increaseFreq(node)
		return
	}

	if c.size == c.capacity {
		l := c.lists[c.minFreq]
		oldTail := l.removeTail()
		delete(c.nodes, oldTail.key)
		c.size--
	}

	node := &Node{
		key:  key,
		val:  value,
		freq: 1,
	}
	c.nodes[key] = node
	c.getList(1).addToHead(node)
	c.minFreq = 1
	c.size++
}

func (c *LFUCache) increaseFreq(node *Node) {
	oldList := c.lists[node.freq]
	oldList.remove(node)
	if node.freq == c.minFreq && oldList.size == 0 {
		delete(c.lists, c.minFreq)
		c.minFreq++
	}

	node.freq++

	l := c.getList(node.freq)
	l.addToHead(node)
}

func (c *LFUCache) getList(freq int) *DList {
	if l, ok := c.lists[freq]; ok {
		return l
	}

	l := newDList()
	c.lists[freq] = l
	return l
}

func (l *DList) removeTail() *Node {
	oldTail := l.tail.prev
	l.remove(oldTail)
	return oldTail
}

func (l *DList) addToHead(node *Node) {
	l.remove(node)

	oldHead := l.head.next
	l.head.next = node
	oldHead.prev = node

	node.prev = l.head
	node.next = oldHead

	l.size++
}

func newDList() *DList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head

	l := &DList{
		head: head,
		tail: tail,
	}
	return l
}

func (l *DList) remove(node *Node) {
	if node == nil || node.prev == nil || node.next == nil {
		return
	}

	np, nn := node.prev, node.next
	np.next, nn.prev = nn, np

	node.prev, node.next = nil, nil

	l.size--
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
