package all_oone_data_struct

type queue struct {
	pre, next *queue
	members   map[string]*item
}

type item struct {
	val   string
	count int
	q     *queue
}

type AllOne struct {
	m          map[string]*item
	head, tail *queue
}

func Constructor() AllOne {
	return AllOne{
		m: make(map[string]*item),
	}
}

func (o *AllOne) Inc(key string) {
	if _, exists := o.m[key]; !exists {
		o.m[key] = &item{
			val: key,
		}
	}

	one := o.m[key]
	one.count++

}

func (o *AllOne) Dec(key string) {
	one := o.m[key]
	one.count--

}

func (o *AllOne) GetMaxKey() string {

}

func (o *AllOne) GetMinKey() string {

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
