package min_stack

type MinStackNES struct {
	mv  int
	stk []int
}

func ConstructorNES() MinStackNES {
	return MinStackNES{
		mv:  -1,
		stk: make([]int, 0),
	}
}

func (this *MinStackNES) Push(val int) {
	if len(this.stk) == 0 {
		this.stk = append(this.stk, 0)
		this.mv = val
	} else {
		diff := val - this.mv
		this.stk = append(this.stk, diff)
		if diff < 0 {
			this.mv = val
		}
	}
}

func (this *MinStackNES) Pop() {
	n := len(this.stk)
	diff := this.stk[n-1]
	var top int
	if diff < 0 {
		top = this.mv
		this.mv = top - diff
	} else {
		top = this.mv + diff
	}
	return top
}

func (this *MinStackNES) Top() int {
}

func (this *MinStackNES) GetMin() int {
}
