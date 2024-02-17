package longest_univalue_path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(n *TreeNode) int
	var ans int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		var l, r int
		var ll, rl int
		if n.Left != nil {
			ll = dfs(n.Left)
			if n.Left.Val == n.Val {
				l = ll + 1
			}
		}
		if n.Right != nil {
			rl = dfs(n.Right)
			if n.Right.Val == n.Val {
				r = rl + 1
			}
		}
		ta := maxInt(maxInt(ll, rl), l+r)
		ans = maxInt(ans, ta)
		return maxInt(l, r)
	}
	dfs(root)

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
