package lru_cache

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// 看上去span记录所有有效距离的方法是不可行的啊！在put新增的情况下，无法在O(1)找到最后一个有效的
// type loopQueue struct {
//	items      []int
//	size       int
//	head, tail int
//}
//
//func NewLoopQueue(n int) *loopQueue {
//	return &loopQueue{
//		items: make([]int, n),
//		size:  n,
//		head:  0,
//		tail:  0,
//	}
//}
//
//func (l *loopQueue) push(x int) (int, bool) {
//	var ok bool
//	var y int
//	if l.head-l.tail >= l.size {
//		y = l.items[l.tail%l.size]
//		ok = true
//	}
//	l.items[l.head%l.size] = x
//	l.head++
//	l.tail++
//	return y, ok
//}
//
//type cacheItem struct {
//	value int
//	index int
//}
//
//func NewCacheItem(v, i int) *cacheItem {
//	return &cacheItem{
//		value: v,
//		index: i,
//	}
//}
//
//type LRUCache struct {
//	m       map[int]*cacheItem
//	counter int
//	span    int
//	cap     int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		m:    make(map[int]*cacheItem),
//		span: capacity,
//		cap:  capacity,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	if x, ok := this.m[key]; ok {
//		if x.index < this.counter-this.span {
//			delete(this.m, key)
//			return -1
//		}
//		x.index = this.counter
//		this.counter++
//		this.span++
//		return x.value
//	}
//	return -1
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if x, ok := this.m[key]; ok {
//		if x.index < this.counter-this.span {
//			if this.span > this.cap {
//				this.span--
//			}
//		} else {
//			this.span++
//		}
//		x.index = this.counter
//		x.value = value
//		this.counter++
//		return
//	}
//	if this.span > this.cap {
//		// 如果一个数get n次，就会缔造出n次重复的，但却只减去一次
//		this.span--
//	}
//	x := NewCacheItem(value, this.counter)
//	this.m[key] = x
//	this.counter++
//}
