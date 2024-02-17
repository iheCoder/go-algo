package lowest_common_ancestor_of_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如果想比较得到LCA，那么p、q必须放在一起，可放在一起必定会有很多不必要的比较
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	if (p.Val-root.Val)^(q.Val-root.Val) < 0 {
		return root
	}

	if p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return lowestCommonAncestor(root.Left, p, q)
}

//func lca(root, p *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	if root == p {
//		return root
//	}
//
//	var pp *TreeNode
//	if p.Val < root.Val && root.Left != nil {
//		pp = lca(root.Left, p)
//	} else if p.Val > root.Val && root.Right != nil {
//		pp = lca(root.Right, p)
//	}
//
//	return pp
//}

// 	var pp, qp *TreeNode
//	if p != nil {
//		if p.Val < root.Val && root.Left != nil {
//			pp = lowestCommonAncestor(root.Left, p, nil)
//		} else if p.Val > root.Val && root.Right != nil {
//			pp = lowestCommonAncestor(root.Right, p, nil)
//		}
//	}
//
//	if q != nil {
//		if q.Val < root.Val && root.Left != nil {
//			qp = lowestCommonAncestor(root.Left, nil, q)
//		} else if q.Val > root.Val && root.Right != nil {
//			qp = lowestCommonAncestor(root.Right, nil, q)
//		}
//	}
//
//	if pp != nil && qp != nil && pp != qp {
//		return root
//	}
//	if pp != nil {
//		return pp
//	}
