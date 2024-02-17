package min_stack

type item struct {
	val int
	mv  int
}

type MinStack struct {
	items []*item
}

func Constructor() MinStack {
	return MinStack{items: make([]*item, 0)}
}

func (this *MinStack) Push(val int) {
	mv := val
	n := len(this.items)
	if n != 0 && this.items[n-1].mv < val {
		mv = this.items[n-1].mv
	}
	this.items = append(this.items, &item{mv: mv, val: val})
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	n := len(this.items)
	return this.items[n-1].val
}

func (this *MinStack) GetMin() int {
	n := len(this.items)
	return this.items[n-1].mv
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
