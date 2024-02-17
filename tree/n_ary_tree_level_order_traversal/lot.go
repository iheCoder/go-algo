package n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	stk := []*Node{root}
	ans := make([][]int, 0)
	var n int
	for len(stk) > 0 {
		n = len(stk)
		item := make([]int, 0, n)
		children := make([]*Node, 0)
		for _, s := range stk {
			item = append(item, s.Val)
			children = append(children, s.Children...)
		}
		stk = children
		ans = append(ans, item)
	}
	return ans
}
