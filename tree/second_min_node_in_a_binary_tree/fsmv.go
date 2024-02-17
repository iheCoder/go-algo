package second_min_node_in_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil || root.Right == nil {
		return -1
	}

	var lmin, rmin int
	if root.Left.Val > root.Val {
		lmin = root.Left.Val
		rmin = findSecondMinimumValue(root.Right)
	} else if root.Right.Val > root.Val {
		lmin = findSecondMinimumValue(root.Left)
		rmin = root.Right.Val
	} else {
		lmin = findSecondMinimumValue(root.Left)
		rmin = findSecondMinimumValue(root.Right)
	}

	return minInt(lmin, rmin)
}

func minInt(x, y int) int {
	if x == -1 {
		return y
	}
	if y == -1 {
		return x
	}
	if x < y {
		return x
	}
	return y
}
