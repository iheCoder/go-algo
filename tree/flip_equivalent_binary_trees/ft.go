package flip_equivalent_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == nil && root2 == nil
	}

	if root1.Val != root2.Val {
		return false
	}

	l1, l2, r1, r2 := -1, -1, -1, -1
	if root1.Left != nil {
		l1 = root1.Left.Val
	}
	if root1.Right != nil {
		r1 = root1.Right.Val
	}
	if root2.Left != nil {
		l2 = root2.Left.Val
	}
	if root2.Right != nil {
		r2 = root2.Right.Val
	}

	if l1*l2*r1*r2 < 0 {
		return false
	}

	if l1 == l2 {
		if r1 != r2 {
			return false
		}

		return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
	} else if l1 == r2 {
		if r1 != l2 {
			return false
		}

		return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
	}

	return false
}
