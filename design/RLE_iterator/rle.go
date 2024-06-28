package RLE_iterator

type item struct {
	val   int
	count int
}

type RLEIterator struct {
	total int
	items []*item
}

func Constructor(encoding []int) RLEIterator {
	var total int
	n := len(encoding)
	items := make([]*item, 0, n)
	for i := 0; i < n; i += 2 {
		if encoding[i] == 0 {
			continue
		}

		im := &item{
			val:   encoding[i+1],
			count: encoding[i],
		}
		total += encoding[i]
		items = append(items, im)
	}

	return RLEIterator{
		total: total,
		items: items,
	}
}

func (rle *RLEIterator) Next(n int) int {
	if n >= rle.total {
		last := -1
		if n == rle.total {
			last = rle.items[len(rle.items)-1].val
		}
		rle.total = 0
		return last
	}

	rle.total -= n

	var i int
	last := -1
	for n > 0 {
		for rle.items[i].count <= n {
			last = rle.items[i].val
			n -= rle.items[i].count
			i++
		}

		if n <= 0 {
			break
		}

		rle.items[i].count -= n
		n = 0
		last = rle.items[i].val
	}

	rle.items = rle.items[i:]
	return last
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */
