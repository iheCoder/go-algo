package cousins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	dx, dy := -1, -1
	ans := true
	var f func(n *TreeNode, d int) bool
	f = func(n *TreeNode, d int) bool {
		if n == nil {
			return false
		}

		var found bool
		if n.Val == x {
			dx = d
			found = true
		}
		if n.Val == y {
			dy = d
			found = true
		}
		var lfound, rfound bool
		lfound = f(n.Left, d+1)
		rfound = f(n.Right, d+1)
		if lfound && rfound && dx == dy && dx == d+1 {
			ans = false
			return true
		}
		return found
	}
	f(root, 0)
	if !ans {
		return false
	}
	return dx == dy
}
