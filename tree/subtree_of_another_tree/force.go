package subtree_of_another_tree

func isSubtreeForce(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	return check(s, t) || isSubtreeForce(s.Left, t) || isSubtreeForce(s.Right, t)
}

func check(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	if r1.Val == r2.Val {
		return check(r1.Left, r2.Left) && check(r1.Right, r2.Right)
	}

	return false
}
