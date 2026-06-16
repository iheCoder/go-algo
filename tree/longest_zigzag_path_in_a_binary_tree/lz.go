package longest_zigzag_path_in_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	var ans int
	var zz func(n *TreeNode, isLeft bool) int
	zz = func(n *TreeNode, isLeft bool) int {
		if n == nil {
			return 0
		}

		l := zz(n.Left, true) + 1
		r := zz(n.Right, false) + 1
		ans = maxInt(ans, maxInt(l, r))

		if isLeft {
			return r
		} else {
			return l
		}
	}
	zz(root, true)
	zz(root, false)

	return ans - 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
