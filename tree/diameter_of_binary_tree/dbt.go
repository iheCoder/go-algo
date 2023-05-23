package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = 1

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dbt(root)
	defer func() {
		ans = 1
	}()
	return ans - 1
}

func dbt(n *TreeNode) int {
	if n == nil {
		return 0
	}
	l := dbt(n.Left)
	r := dbt(n.Right)
	ans = maxInt(ans, l+r+1)
	return maxInt(l, r) + 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
