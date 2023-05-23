package del_nodes_and_return_forest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	dm := make(map[int]bool)
	for _, td := range to_delete {
		dm[td] = true
	}

	var ans []*TreeNode
	var f func(n *TreeNode) *TreeNode
	f = func(n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}

		l := f(n.Left)
		r := f(n.Right)

		if dm[n.Val] {
			if l != nil {
				ans = append(ans, l)
			}
			if r != nil {
				ans = append(ans, r)
			}
			return nil
		}
		n.Left = l
		n.Right = r
		return n
	}
	if f(root) != nil {
		ans = append(ans, root)
	}

	return ans
}
