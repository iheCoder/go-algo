package peeking_iterator

//   Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	items []int
	i     int
}

func Constructor(iter *Iterator) *PeekingIterator {
	items := make([]int, 0)
	for iter.hasNext() {
		items = append(items, iter.next())
	}
	return &PeekingIterator{
		items: items,
	}
}

func (pi *PeekingIterator) hasNext() bool {
	if pi.i < len(pi.items) {
		return true
	}
	return false
}

func (pi *PeekingIterator) next() int {
	x := pi.items[pi.i]
	pi.i++
	return x
}

func (pi *PeekingIterator) peek() int {
	return pi.items[pi.i]
}
