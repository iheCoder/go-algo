package design_linked_list

type bdl struct {
	prev, next *bdl
	val        int
}

type MyLinkedList struct {
	head, tail *bdl
	size       int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	p := this.findItemByIndex(index)
	if p == nil {
		return -1
	}
	return p.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	p := &bdl{
		next: this.head,
		val:  val,
	}
	if this.head != nil {
		this.head.prev = p
	} else {
		this.tail = p
	}
	this.head = p
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	p := &bdl{
		prev: this.tail,
		val:  val,
	}
	if this.tail != nil {
		this.tail.next = p
	} else {
		this.head = p
	}
	this.tail = p
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	} else if index == 0 {
		this.AddAtHead(val)
		return
	}

	ip := &bdl{val: val}
	p := this.findItemByIndex(index)
	ip.prev = p.prev
	p.prev.next = ip
	ip.next = p
	p.prev = ip
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	p := this.findItemByIndex(index)
	if p.prev != nil {
		p.prev.next = p.next
	} else {
		this.head = this.head.next
	}

	if p.next != nil {
		p.next.prev = p.prev
	} else {
		this.tail = this.tail.prev
	}
	this.size--
}

func (this *MyLinkedList) findItemByIndex(index int) *bdl {
	p := this.head
	var i int
	for p != nil && i < index {
		p = p.next
		i++
	}
	return p
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
