package segment_tree

const emptyLazyNode = 0

type SegmentTree struct {
	// 原数组
	arr []int
	// 线段树
	st []int
	// 懒惰标记。延迟对节点修改，在访问时才进行修改
	lazyMark []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	if len(arr) == 0 {
		return nil
	}

	segTree := SegmentTree{
		arr:      arr,
		st:       make([]int, 4*len(arr)),
		lazyMark: make([]int, 4*len(arr)),
	}

	for i := range segTree.lazyMark {
		//fill LazyTree with empty lazy nodes
		segTree.lazyMark[i] = emptyLazyNode
	}

	//starts with node 1 and interval [0, len(arr)-1] inclusive
	segTree.Build(1, 0, len(arr)-1)

	return &segTree
}

func (s *SegmentTree) Build(node int, left, right int) {
	// 当左右边界相等时，线段节点就是数组节点
	if left == right {
		s.st[node] = s.arr[left]
		return
	}

	// 将数组分为两部分进行构建
	mid := (left + right) / 2

	s.Build(2*node, left, mid)
	s.Build(2*node, mid+1, right)

	s.st[node] = s.st[2*node] + s.st[2*node+1]
}

func (s *SegmentTree) Update(node int, left, right int, ql, qr int, val int) {
	// 触发懒惰标记更新
	s.Propagate(node, left, right)

	// 无效值返回
	if ql > qr || left > right {
		return
	}

	if left >= ql && right <= qr {
		// 当前区间为查询区间的子区间，则更新懒惰标记，并传递更新
		s.lazyMark[node] += val
		s.Propagate(node, left, right)
	} else {
		// 分左右子区间递归更新
		mid := (left + right) / 2

		s.Update(2*node, left, mid, ql, minInt(mid, qr), val)
		s.Update(2*node+1, mid+1, right, maxInt(ql, mid+1), qr, val)

		s.st[node] = s.st[2*node] + s.st[2*node+1]
	}
}

func (s *SegmentTree) Propagate(node, left, right int) {
	// 如果存在标记
	if s.lazyMark[node] != emptyLazyNode {
		s.st[node] += (right - left + 1) * s.lazyMark[node]

		if left == right {
			// 被标记的是叶节点，那么直接更新
			s.arr[left] += s.lazyMark[node]
		} else {
			// 其他情况下，更新左右子线段节点的标记
			s.lazyMark[2*node] += s.lazyMark[node]
			s.lazyMark[2*node+1] += s.lazyMark[node]
		}

		s.lazyMark[node] = emptyLazyNode
	}
}

// left、right指向节点区间
// ql、qr指向查询区间
func (s *SegmentTree) Query(node int, left, right int, ql, qr int) int {
	// 无效值直接返回
	if (ql > qr) || (left > right) {
		return 0
	}

	// 若存在标记，则更新
	s.Propagate(node, left, right)

	// 当前区间为询问区间的子集时直接返回当前区间的和
	if left >= ql && right <= qr {
		return s.st[node]
	}

	// 继续分左右子区间递归查询
	mid := (left + right) / 2
	leftNodeSum := s.Query(2*node, left, mid, ql, minInt(mid, qr))
	rightNodeSum := s.Query(2*node+1, mid+1, right, maxInt(ql, mid+1), qr)

	return leftNodeSum + rightNodeSum
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
