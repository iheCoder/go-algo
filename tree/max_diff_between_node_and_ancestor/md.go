package max_diff_between_node_and_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxAncestorDiffHelper(root, root.Val, root.Val)
}

func maxAncestorDiffHelper(root *TreeNode, min, max int) int {
	if root == nil {
		return max - min
	}

	min = minInt(min, root.Val)
	max = maxInt(max, root.Val)

	return maxInt(maxAncestorDiffHelper(root.Left, min, max), maxAncestorDiffHelper(root.Right, min, max))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
