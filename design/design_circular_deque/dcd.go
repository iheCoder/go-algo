package design_circular_deque

type MyCircularDeque struct {
	items        []int
	size, length int
	// tail 指向下一个尾部
	// head 指向下一个头部
	head, tail int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		items:  make([]int, k),
		size:   k,
		length: 0,
		head:   0,
		tail:   0,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}

	if d.head == -1 {
		d.head = d.size - 1
	}

	d.items[d.head] = value
	d.length++

	if d.head == d.tail && !d.IsFull() {
		d.tail = d.head + 1
	}
	d.head--
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}

	if d.tail == d.size {
		d.tail = 0
	}

	d.items[d.tail] = value
	d.length++

	if d.tail == d.head && !d.IsFull() {
		d.head = d.tail - 1
	}
	d.tail++

	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}

	d.head = d.increase(d.head)
	d.length--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}

	d.tail = d.decrease(d.tail)
	d.length--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.items[d.increase(d.head)]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.items[d.decrease(d.tail)]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.length == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.length == d.size
}

func (d *MyCircularDeque) increase(x int) int {
	x++
	if x == d.size {
		x = 0
	}
	return x
}

func (d *MyCircularDeque) decrease(x int) int {
	x--
	if x == -1 {
		x = d.size - 1
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
