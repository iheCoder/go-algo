package check_completeness_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stk := []*TreeNode{root}
	var p *TreeNode
	var lastEmpty bool
	for len(stk) != 0 {
		p = stk[0]
		stk = stk[1:]
		if p.Left == nil && p.Right != nil {
			return false
		}
		if p.Left == nil && p.Right == nil {
			lastEmpty = true
			continue
		}
		if lastEmpty {
			return false
		}
		if p.Left != nil && p.Right != nil {
			stk = append(stk, p.Left)
			stk = append(stk, p.Right)
			continue
		}
		lastEmpty = true
		stk = append(stk, p.Left)
	}

	return true
}
