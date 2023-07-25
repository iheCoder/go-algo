package impl_queue_using_stacks

type stack struct {
	items []int
}

func (s *stack) push(x int) {
	s.items = append(s.items, x)
}

func (s *stack) pop() int {
	x := s.peek()
	s.items = s.items[:len(s.items)-1]
	return x
}

func (s *stack) peek() int {
	return s.items[len(s.items)-1]
}

func (s *stack) size() int {
	return len(s.items)
}

type MyQueue struct {
	inStk  stack
	outStk stack
	size   int
}

func Constructor() MyQueue {
	return MyQueue{
		inStk:  stack{items: make([]int, 0)},
		outStk: stack{items: make([]int, 0)},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStk.push(x)
	this.size++
}

func (this *MyQueue) Pop() int {
	if this.outStk.size() == 0 {
		for this.inStk.size() > 0 {
			this.outStk.push(this.inStk.pop())
		}
	}

	this.size--
	return this.outStk.pop()
}

func (this *MyQueue) Peek() int {
	if this.outStk.size() == 0 {
		for this.inStk.size() > 0 {
			this.outStk.push(this.inStk.pop())
		}
	}

	return this.outStk.peek()
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
