package design_circular_queue

type MyCircularQueue struct {
	l          int
	items      []int
	head, tail int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		items: make([]int, k),
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	q.items[q.tail%len(q.items)] = value
	q.tail++
	q.l++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.head++
	q.l--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.items[q.head%len(q.items)]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.items[(q.tail-1)%len(q.items)]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.l == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.l == len(q.items)
}
