package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	var r []int
	p := root
	var prev *TreeNode
	for len(stack) != 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		// p.Right == prev就是防止重复访问
		if p.Right == nil || p.Right == prev {
			stack = stack[:len(stack)-1]
			r = append(r, p.Val)
			prev = p
			p = nil
		} else {
			p = p.Right
		}
	}

	return r
}
