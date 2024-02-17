package lru_cache

import "testing"

func TestLC(t *testing.T) {
	lc := Constructor(2)
	lc.Put(1, 1)
	lc.Put(2, 2)
	t.Log(lc.Get(1))
	lc.Put(3, 3)
	t.Log(lc.Get(2))
	lc.Put(4, 4)
	t.Log(lc.Get(1))
	t.Log(lc.Get(3))
	t.Log(lc.Get(4))
	lc.Put(3, 33)
	lc.Put(5, 5)
	t.Log(lc.Get(4))
	t.Log(lc.Get(3))

	//t.Log(lc.Get(2))
	//lc.Put(2,6)
	//t.Log(lc.Get(1))
	//lc.Put(1,5)
	//lc.Put(1,2)
	//t.Log(lc.Get(1))
	//t.Log(lc.Get(2))
}
