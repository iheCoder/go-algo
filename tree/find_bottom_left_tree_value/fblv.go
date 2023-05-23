package find_bottom_left_tree_value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftmost := root.Val
	stack := []*TreeNode{root}
	var p *TreeNode
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
		if vc == 0 && len(stack) != 0 {
			vc = len(stack)
			leftmost = stack[0].Val
		}
	}

	return leftmost
}
