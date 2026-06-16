package design_a_stack_with_increment_op

type CustomStack struct {
	arr     []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		arr:     make([]int, 0, maxSize),
		maxSize: maxSize,
	}
}

func (cs *CustomStack) Push(x int) {
	if len(cs.arr) >= cs.maxSize {
		return
	}

	cs.arr = append(cs.arr, x)
}

func (cs *CustomStack) Pop() int {
	if len(cs.arr) == 0 {
		return -1
	}

	x := cs.arr[len(cs.arr)-1]
	cs.arr = cs.arr[:len(cs.arr)-1]
	return x
}

func (cs *CustomStack) Increment(k int, val int) {
	for i := 0; i < k; i++ {
		if i >= len(cs.arr) {
			break
		}

		cs.arr[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
