package max_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return md(root, 0)
}

func md(p *TreeNode, d int) int {
	if p == nil {
		return d
	}
	d++
	d = maxInt(md(p.Left, d), md(p.Right, d))
	return d
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
