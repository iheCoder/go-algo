package house_robber_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r := hr(root)
	return maxInt(r[0], r[1])
}

func hr(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val, 0}
	}

	l := hr(root.Left)
	r := hr(root.Right)

	selected := root.Val + l[1] + r[1]
	unselect := maxInt(l[0], l[1]) + maxInt(r[0], r[1])
	return []int{selected, unselect}
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
