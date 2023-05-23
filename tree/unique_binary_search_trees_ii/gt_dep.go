package unique_binary_search_trees_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 始终无法得到[2,1,3]这样的binary tree，并且题目要求是binary search tree
func generateTreesDep(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{
			{
				Val: n,
			},
		}
	}

	tn := generateTreesDep(n - 1)

	r := make([]*TreeNode, 0)
	for _, node := range tn {
		rl := &TreeNode{
			Val:  n,
			Left: node,
		}
		rr := &TreeNode{
			Val:   n,
			Right: node,
		}
		r = append(r, rl)
		r = append(r, rr)
	}

	return r
}

// 	if n == 2 {
//		l := &TreeNode{
//			Val: n,
//		}
//		r := &TreeNode{
//			Val: n,
//		}
//
//		rl := &TreeNode{
//			Val:  n - 1,
//			Left: l,
//		}
//		rr := &TreeNode{
//			Val:   n - 1,
//			Right: r,
//		}
//		return []*TreeNode{
//			rl,
//			rr,
//		}
//	}
