package max_product_of_splitted_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const maxVal = 1000000007

func maxProduct(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var subTreeSum []int
	var f func(n *TreeNode) int
	f = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		if n.Left == nil && n.Right == nil {
			subTreeSum = append(subTreeSum, n.Val)
			return n.Val
		}
		sts := f(n.Left) + f(n.Right) + n.Val
		subTreeSum = append(subTreeSum, sts)
		return sts
	}

	sum := f(root)
	var r int
	for _, s := range subTreeSum {
		r = maxInt(r, s*(sum-s))
	}
	if r > maxVal {
		r %= maxVal
	}
	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
