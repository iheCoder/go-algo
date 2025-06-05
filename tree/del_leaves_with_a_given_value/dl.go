package del_leaves_with_a_given_value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var del func(n *TreeNode) bool
	del = func(n *TreeNode) bool {
		if n == nil {
			return true
		}

		if n.Left == nil && n.Right == nil {
			return n.Val == target
		}

		ld, rd := del(n.Left), del(n.Right)
		if ld {
			n.Left = nil
		}
		if rd {
			n.Right = nil
		}

		if n.Val == target && ld && rd {
			return true
		}

		return false
	}
	if del(root) {
		return nil
	}

	return root
}
