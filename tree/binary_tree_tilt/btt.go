package binary_tree_tilt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r, _ := ft(root)
	return r
}

func ft(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	switch {
	case root.Left == nil && root.Right == nil:
		return 0, root.Val
	}

	lds, ls := ft(root.Left)
	rds, rs := ft(root.Right)

	return lds + rds + absDiff(ls-rs), ls + rs + root.Val
}

func absDiff(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
