package range_sum_of_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var sum int
	var ls, rs int
	if root.Val >= low && root.Val <= high {
		sum += root.Val
		ls = rangeSumBST(root.Left, low, high)
		rs = rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		ls = rangeSumBST(root.Left, low, high)
	} else {
		rs = rangeSumBST(root.Right, low, high)
	}
	sum += ls + rs
	return sum
}
