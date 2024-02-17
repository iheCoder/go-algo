package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	verCnt := 1
	var p *TreeNode
	var r []int
	for len(stack) != 0 {
		p = stack[0]
		stack = stack[1:]
		verCnt--

		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}

		if verCnt == 0 {
			verCnt = len(stack)
			r = append(r, p.Val)
		}
	}

	return r
}
