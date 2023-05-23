package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(height(root.Left)-height(root.Right)) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxInt(height(root.Left), height(root.Right)) + 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
