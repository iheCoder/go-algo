package exam_room

type linkedItem struct {
	val  int
	next *linkedItem
}

type ExamRoomDep struct {
	n          int
	head, tail *linkedItem
}

//func Constructor(n int) ExamRoomDep {
//	return ExamRoomDep{n: n}
//}

func (er *ExamRoomDep) Seat() int {
	if er.head == nil {
		item := &linkedItem{
			val: 0,
		}
		er.head = item
		er.tail = item
		return item.val
	}

	if er.head.next == nil {
		item := &linkedItem{
			val: er.n - 1,
		}
		er.head.next = item
		er.tail = item
		return item.val
	}

	if er.head.val != 0 {
		item := &linkedItem{
			val:  0,
			next: er.head,
		}
		er.head = item
		return item.val
	}

	if er.tail.val != er.n-1 {
		item := &linkedItem{
			val: er.n - 1,
		}
		er.tail.next = item
		er.tail = item
		return item.val
	}

	var maxDis int
	var val int
	cur, next := er.head, er.head.next
	var mprev, mnext *linkedItem
	for next != nil {
		diff := next.val - cur.val
		if diff > maxDis {
			val = diff / 2
			maxDis = diff
			mprev, mnext = cur, next
		}
		cur = next
		next = cur.next
	}
	item := &linkedItem{
		val:  val,
		next: mnext,
	}
	mprev.next = item
	return item.val
}

func (er *ExamRoomDep) Leave(val int) {
	if er.head.val == val {
		er.head = er.head.next
		return
	}

	prev := er.head
	next := er.head.next
	for next != nil {
		if next.val == val {
			prev.next = next.next
			return
		}
		prev = next
		next = prev.next
	}
}

/**
 * Your ExamRoomDep object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
