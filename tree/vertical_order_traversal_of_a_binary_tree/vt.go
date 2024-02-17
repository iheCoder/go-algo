package vertical_order_traversal_of_a_binary_tree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeHeapItem struct {
	col int
	val int
}

type nodeHeap []*nodeHeapItem

func (n *nodeHeap) Len() int {
	return len(*n)
}

func (n *nodeHeap) Less(i, j int) bool {
	ii, ji := (*n)[i], (*n)[j]
	if ii.col != ji.col {
		return ii.col < ji.col
	}
	return ii.val < ji.val
}

func (n *nodeHeap) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n *nodeHeap) Push(x interface{}) {
	*n = append(*n, x.(*nodeHeapItem))
}

func (n *nodeHeap) Pop() interface{} {
	l := len(*n)
	x := (*n)[l-1]
	*n = (*n)[:l-1]
	return x
}

func (n *nodeHeap) list() []int {
	if n.Len() == 0 {
		return nil
	}

	sort.Sort(n)

	l := len(*n)
	r := make([]int, l)
	for i := 0; i < l; i++ {
		r[i] = (*n)[i].val
	}
	return r
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	m := make(map[int]*nodeHeap)
	minRow := 0
	var f func(n *TreeNode, row, col int)
	f = func(n *TreeNode, row, col int) {
		if n == nil {
			return
		}
		if m[row] == nil {
			nhi := (nodeHeap)(make([]*nodeHeapItem, 0))
			m[row] = &nhi
		}
		m[row].Push(&nodeHeapItem{
			col: col,
			val: n.Val,
		})
		minRow = minInt(minRow, row)
		f(n.Left, row-1, col+1)
		f(n.Right, row+1, col+1)
	}
	f(root, 0, 0)
	r := make([][]int, 0, len(m))
	var h *nodeHeap
	var ok bool
	for i := minRow; ; i++ {
		if h, ok = m[i]; !ok {
			return r
		}
		r = append(r, h.list())
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
