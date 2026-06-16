package lru_cache

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	m              map[int]*item
	head, tail     *item
	capacity, size int
}

type item struct {
	key, val   int
	next, prev *item
}

func Constructor(capacity int) LRUCache {
	head, tail := &item{}, &item{}
	head.next, tail.prev = tail, head

	lru := LRUCache{
		m:        make(map[int]*item),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}

	return lru
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.m[key]; ok {
		c.moveToHead(key)
		return v.val
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.m[key]; ok {
		c.m[key].val = value
		c.moveToHead(key)
		return
	}

	oldHead := c.head.next
	it := &item{
		key:  key,
		val:  value,
		next: oldHead,
		prev: c.head,
	}
	c.head.next = it

	oldHead.prev = it
	it.next = oldHead

	c.m[key] = it
	c.size++

	if c.size > c.capacity {
		c.removeTail()
		return
	}
}

func (c *LRUCache) removeTail() {
	oldTail := c.tail.prev
	newTail := oldTail.prev
	if newTail != nil {
		newTail.next = c.tail
		c.tail.prev = newTail
	}

	delete(c.m, oldTail.key)
	oldTail = nil
	c.size--
}

func (c *LRUCache) moveToHead(key int) {
	it := c.m[key]

	oldHead := c.head.next
	if oldHead == it {
		return
	}

	prev := it.prev
	next := it.next

	prev.next = next
	next.prev = prev

	oldHead.prev = it
	it.next = oldHead

	it.prev = c.head
	c.head.next = it

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
