package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtreeDep(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil && root == nil {
		return true
	}
	if subRoot == nil || root == nil {
		return false
	}

	if root.Val == subRoot.Val {
		return isSubtreeDep(root.Left, subRoot.Left) && isSubtreeDep(root.Right, subRoot.Right)
	}

	return isSubtreeDep(root.Left, subRoot) || isSubtreeDep(root.Right, subRoot)
}
