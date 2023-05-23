package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 还别说直接判断root是否等于p、q比在p、q为答案的情况下是有效的
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil || right == nil {
		if left != nil {
			return left
		}
		if right != nil {
			return right
		}
		return nil
	}

	return root
}
