package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil || root.Right == nil {
		if root.Left == nil && root.Right == nil {
			return true
		}
		return false
	}

	return isSym(root.Left, root.Right)
}

func isSym(left, right *TreeNode) bool {
	if left == nil || right == nil {
		if left == nil && right == nil {
			return true
		}
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if !isSym(left.Left, right.Right) {
		return false
	}
	if !isSym(left.Right, right.Left) {
		return false
	}

	return true
}
