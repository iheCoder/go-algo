package increasing_order_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stk := make([]*TreeNode, 0)
	p := root
	dr := &TreeNode{}
	last := dr
	for p != nil || len(stk) != 0 {
		for p != nil {
			stk = append(stk, p)
			p = p.Left
		}
		p = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		last.Right = p
		last = p
		last.Left = nil
		p = p.Right
	}

	return dr.Right
}
