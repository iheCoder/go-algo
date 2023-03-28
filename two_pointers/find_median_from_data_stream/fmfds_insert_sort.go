package find_median_from_data_stream

type listNode struct {
	next *listNode
	val  int
}

type MedianFinderInsertSort struct {
	arr   []int
	head  *listNode
	count int
}

func ConstructorIS() MedianFinderInsertSort {
	//return MedianFinderInsertSort{arr: make([]int, 0)}
	return MedianFinderInsertSort{}
}

func (this *MedianFinderInsertSort) AddNum(num int) {
	//this.arr = append(this.arr, num)
	this.insertSortedNums(num)
}

func (this *MedianFinderInsertSort) FindMedian() float64 {
	return this.findMedian()
}

func (this *MedianFinderInsertSort) findMedian() float64 {
	if this.count == 0 {
		return 0
	}
	medIndex := make([]int, 0)
	if this.count%2 == 1 {
		medIndex = append(medIndex, this.count/2)
	} else {
		medIndex = append(medIndex, this.count/2-1, this.count/2)
	}

	var sum float64
	p := this.head.next
	var i int
	for _, index := range medIndex {
		for ; p != nil; p = p.next {
			if i == index {
				sum += float64(p.val)
				break
			}
			i++
		}
	}

	return sum / float64(len(medIndex))
}

func (this *MedianFinderInsertSort) insertSortedNums(num int) {
	this.count++

	n := &listNode{
		next: nil,
		val:  num,
	}
	if this.head == nil {
		this.head = &listNode{
			next: n,
		}
		return
	}

	prev := this.head
	p := this.head.next
	for p != nil {
		if p.val > num {
			break
		}
		prev = p
		p = p.next
	}

	prev.next = n
	n.next = p
}

/**
 * Your MedianFinderInsertSort object will be instantiated and called as such:
 * obj := ConstructorIS();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
