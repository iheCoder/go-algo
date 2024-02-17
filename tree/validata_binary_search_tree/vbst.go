package validata_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	r, _, _ := ivbst(root)
	return r
}

func ivbst(root *TreeNode) (bool, *int, *int) {
	if root == nil {
		return true, nil, nil
	}

	minVal, maxVal := &(root.Val), &(root.Val)
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false, nil, nil
		}
		minVal = &(root.Left.Val)
		ok, mx, mv := ivbst(root.Left)
		if !ok {
			return false, nil, nil
		}
		if mv != nil {
			if *mv >= root.Val {
				return false, nil, nil
			}
		}
		if mx != nil {
			minVal = mx
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false, nil, nil
		}
		maxVal = &(root.Right.Val)
		ok, mv, mx := ivbst(root.Right)
		if !ok {
			return false, nil, nil
		}
		if mv != nil {
			if *mv <= root.Val {
				return false, nil, nil
			}
		}
		if mx != nil {
			maxVal = mx
		}
	}

	return true, minVal, maxVal
}

// func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
//		return false
//	}
//
//	return isValidBST(root.Left) && isValidBST(root.Right)
//}
