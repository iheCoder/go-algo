package binary_tree_coloring_game

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var xNode *TreeNode

	var f func(t *TreeNode) int
	f = func(t *TreeNode) int {
		if t == nil {
			return 0
		}

		if t.Val == x {
			xNode = t
		}

		return 1 + f(t.Left) + f(t.Right)
	}

	f(root)
	lSize, rSize := f(xNode.Left), f(xNode.Right)
	target := (n + 1) / 2
	if lSize >= target || rSize >= target {
		return true
	}

	return n-lSize-rSize-1 >= target
}
