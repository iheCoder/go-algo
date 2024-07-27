package flip_binary_tree_to_match_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	flipped := make([]int, 0)
	var i int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val != voyage[i] {
			flipped = []int{-1}
			return
		}

		i++
		if i < len(voyage) && node.Left != nil && node.Left.Val != voyage[i] {
			flipped = append(flipped, node.Val)
			dfs(node.Right)
			dfs(node.Left)
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)

	if len(flipped) > 0 && flipped[0] == -1 {
		return []int{-1}
	}
	return flipped
}

// 没有考虑到只有左或右节点的情况
//func flipMatchVoyage(root *TreeNode, voyage []int) []int {
//	var ans []int
//
//	findVoyage := func(start, end, target int) int {
//		j := start
//		for j <= end {
//			if target == voyage[j] {
//				return j
//			}
//			j++
//		}
//		return -1
//	}
//
//	var f func(n *TreeNode, i, e int) bool
//	f = func(n *TreeNode, s, e int) bool {
//		if s > e {
//			return true
//		}
//		if n == nil {
//			return false
//		}
//
//		if voyage[s] != n.Val {
//			return false
//		}
//		s++
//
//		var l, r int
//		// 如果当前值等于左节点值，则匹配寻找右节点所在位置
//		if n.Left != nil && n.Left.Val == voyage[s] {
//			rv := n.Right.Val
//			l = s
//			r = findVoyage(s+1, e, rv)
//			if r < 0 {
//				return false
//			}
//		} else if n.Right != nil && n.Right.Val == voyage[s] {
//			lv := n.Left.Val
//			l = findVoyage(s+1, e, lv)
//			r = s
//			if l < 0 {
//				return false
//			}
//
//			ans = append(ans, n.Val)
//		} else {
//			return false
//		}
//
//		if !f(n.Left, l, r) {
//			return false
//		}
//		if !f(n.Right, r, e) {
//			return false
//		}
//
//		return true
//	}
//	if f(root, 0, len(voyage)-1) {
//		return ans
//	}
//
//	return []int{-1}
//}
