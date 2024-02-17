package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hps(root, targetSum, 0)
}

func hps(node *TreeNode, targetSum, sum int) bool {
	if node == nil {
		return false
	}

	sum += node.Val
	if node.Left == nil && node.Right == nil {
		if sum == targetSum {
			return true
		}
		return false
	}

	return hps(node.Left, targetSum, sum) || hps(node.Right, targetSum, sum)
}
