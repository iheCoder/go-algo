package max_depth_of_binary_tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 1)
	stack[0] = root
	var p *TreeNode
	var r int

	vc := 1
	for len(stack) != 0 {
		p = stack[0]
		stack = stack[1:]
		vc--

		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}

		if vc <= 0 {
			vc = len(stack)
			r++
		}
	}

	return r
}
