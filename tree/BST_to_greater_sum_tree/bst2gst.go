package BST_to_greater_sum_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var sum int
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}

		dfs(n.Right)
		sum += n.Val
		n.Val = sum
		dfs(n.Left)
	}
	dfs(root)
	return root
}
