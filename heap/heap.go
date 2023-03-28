package heap

import "sort"

type HeapItem interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

func Pop(h HeapItem) interface{} {
	n := h.Len() - 1
	// 交换堆最顶与最低元素，此时最大堆最上面为最小值
	h.Swap(0, n)
	// 下沉最小值
	down(h, 0, n)
	// 调用堆元素实现的Pop方法，弹出最大元素，并从堆中删除
	return h.Pop()
}

func down(h HeapItem, x int, n int) bool {
	i := x
	// 遍历堆找到最大子节点，若子节点大于插入元素进行交换
	for {
		// 找到i左child，若左child小于0或者大于等于n，返回没有下沉
		child := i*2 + 1
		if child < 0 || child >= n {
			break
		}

		// 找到右child（确保右child没有到最后元素），若右child大于插入元素，则左child与父节点交换。接着插入元素（此时左节点位置）与右节点交换
		rightChild := child + 1
		if rightChild < n && h.Less(i, rightChild) {
			h.Swap(child, i)
			h.Swap(rightChild, i)
			i = rightChild
			continue
		}

		// 若插入元素小于最大子节点，则交换
		if h.Less(i, child) {
			h.Swap(i, child)
			i = child
			// 若插入元素大于等于最大子节点，退出
		} else {
			break
		}
	}
	return i > x
}

func Push(h HeapItem, item interface{}) {
	// 调用堆元素实现的Push方法，将插入元素插入到堆末尾
	h.Push(item)
	up(h, h.Len()-1)
}

func up(h HeapItem, x int) {
	i := x
	for {
		parent := (i - 1) / 2
		if parent == i || h.Less(i, parent) {
			break
		}

		h.Swap(i, parent)
		i = parent
	}
}

// 以上并没有想好最大堆是什么？事实上最大的在最上面其次在左边
