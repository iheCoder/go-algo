package insert_del_getrandom_o1

import "math/rand"

type RandomizedSet struct {
	m     map[int]int
	items []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:     make(map[int]int),
		items: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.items = append(this.items, val)
	this.m[val] = len(this.items) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.m[val]; ok {
		last := len(this.items) - 1
		this.items[i], this.items[last] = this.items[last], this.items[i]
		this.m[this.items[i]] = i
		this.items = this.items[:last]
		delete(this.m, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.items[rand.Intn(len(this.items))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
