package impl_stack_using_queues

type queue struct {
	items []int
}

func NewQueue() *queue {
	return &queue{items: make([]int, 0)}
}

func (q *queue) push(x int) {
	q.items = append(q.items, x)
}

func (q *queue) pop() int {
	x := q.items[0]
	q.items = q.items[1:]
	return x
}

func (q *queue) peek() int {
	return q.items[0]
}

func (q *queue) size() int {
	return len(q.items)
}

type MyStack struct {
	q    *queue
	last int
}

func Constructor() MyStack {
	return MyStack{q: NewQueue()}
}

func (this *MyStack) Push(x int) {
	this.last = x
	this.q.push(x)
}

func (this *MyStack) Pop() int {
	size := this.q.size()
	x := this.last
	y := this.q.pop()
	for i := 0; i < size-1; i++ {
		this.last = y
		this.q.push(y)
		y = this.q.pop()
	}
	return x
}

func (this *MyStack) Top() int {
	return this.last
}

func (this *MyStack) Empty() bool {
	return this.q.size() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
