package insufficient_nodes_in_root_to_leaf_paths

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}

	var f func(node *TreeNode, sum int) bool
	f = func(node *TreeNode, sum int) bool {
		if node == nil {
			return true
		}

		sum += node.Val
		if node.Left == nil && node.Right == nil {
			return sum < limit
		}

		li, ri := f(node.Left, sum), f(node.Right, sum)
		if li {
			node.Left = nil
		}
		if ri {
			node.Right = nil
		}

		return li && ri
	}

	if f(root, 0) {
		return nil
	}
	return root
}
