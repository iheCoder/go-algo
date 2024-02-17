package univalued_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	val := root.Val
	stack := make([]*TreeNode, 0)
	p := root
	var prev *TreeNode
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		if p.Right == nil || p.Right == prev {
			stack = stack[:len(stack)-1]
			if p.Val != val {
				return false
			}
			prev = p
			p = nil
		} else {
			p = p.Right
		}
	}

	return true
}
