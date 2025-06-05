package sum_of_nodes_with_even_valued_grandparent

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var ans int
	var f func(n *TreeNode, gp, p int)
	f = func(n *TreeNode, gp, p int) {
		if n == nil {
			return
		}

		gp--
		if gp == 0 {
			ans += n.Val
		}

		cgp := 0
		if n.Val%2 == 0 {
			cgp = 2
		}
		f(n.Left, p-1, cgp)
		f(n.Right, p-1, cgp)
	}

	f(root, 0, 0)
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
