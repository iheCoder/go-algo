package find_dup_subtrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtreesDep(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	lr := findDuplicateSubtreesDep(root.Left)
	rr := findDuplicateSubtreesDep(root.Right)
	r, _ := fds(root.Left, root.Right)

	lr = append(lr, r...)
	return append(lr, rr...)
}

func merge(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	lr := merge(root.Left)
	rr := merge(root.Right)
	r, ok := fds(root.Left, root.Right)
	if ok {
		r = append(r, root)
	}

	lr = append(lr, r...)
	return append(lr, rr...)
}

func fds(r1, r2 *TreeNode) ([]*TreeNode, bool) {
	if r1 == nil || r2 == nil {
		if r1 == nil && r2 == nil {
			return nil, true
		}
		return nil, false
	}

	if r1.Left == nil || r2.Left == nil || r1.Right == nil || r2.Right == nil {
		if r1.Left == nil && r1.Right == nil && r2.Left == nil && r2.Right == nil && r1.Val == r2.Val {
			return []*TreeNode{r1}, true
		}
		return nil, false
	}

	lns, lok := fds(r1.Left, r2.Left)
	rns, rok := fds(r1.Right, r2.Right)
	if lok && rok {
		lns = append(lns, r1)
		return append(lns, rns...), true
	}
	return append(lns, rns...), false
}
