package smallest_subtree_with_all_the_deepest_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	stk := []*TreeNode{root}
	var tmpStk []*TreeNode

	for len(stk) > 0 {
		for _, node := range stk {
			if node.Left != nil {
				tmpStk = append(tmpStk, node.Left)
			}
			if node.Right != nil {
				tmpStk = append(tmpStk, node.Right)
			}
		}

		if len(tmpStk) == 0 {
			break
		}
		stk = tmpStk
		tmpStk = nil
	}

	if len(stk) == 1 {
		return stk[0]
	}

	m := make(map[*TreeNode]bool)
	for _, node := range stk {
		m[node] = true
	}

	var ans *TreeNode
	var f func(node *TreeNode) bool
	f = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		if m[node] == true {
			return true
		}

		lok, rok := f(node.Left), f(node.Right)

		if !lok && !rok {
			return false
		}

		if lok && rok {
			ans = node
		}

		return true
	}

	f(root)
	return ans
}
