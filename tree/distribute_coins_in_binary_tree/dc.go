package distribute_coins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans int
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := dfs(n.Left), dfs(n.Right)
		ans += absInt(l) + absInt(r)
		return n.Val + l + r - 1
	}
	dfs(root)
	return ans
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
