package min_abs_diff_in_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	r := math.MaxInt
	if root == nil {
		return r
	}

	if root.Left != nil {
		lm := root.Left
		for lm.Right != nil {
			lm = lm.Right
		}
		r = minInt(r, abs(root.Val-lm.Val))
		r = minInt(r, getMinimumDifference(root.Left))
	}
	if root.Right != nil {
		rm := root.Right
		for rm.Left != nil {
			rm = rm.Left
		}
		r = minInt(r, abs(root.Val-rm.Val))
		r = minInt(r, getMinimumDifference(root.Right))
	}

	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
