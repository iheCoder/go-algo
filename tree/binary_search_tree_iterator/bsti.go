package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	arr  []int
	next int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		arr: inorder(root),
	}
}

func inorder(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var r []int
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		r = append(r, p.Val)
		stack = stack[:len(stack)-1]
		p = p.Right
	}

	return r
}

func (this *BSTIterator) Next() int {
	r := this.arr[this.next]
	this.next++
	return r
}

func (this *BSTIterator) HasNext() bool {
	return this.next < len(this.arr)
}
