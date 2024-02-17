package min_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return minInt(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minInt(x, y int) int {
	if x == 0 {
		return y
	} else if y == 0 {
		return x
	}

	if x < y {
		return x
	}
	return y
}
