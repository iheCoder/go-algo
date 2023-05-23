package construct_string_from_binary_tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	r := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return r
	}
	if root.Left != nil {
		r += "(" + tree2str(root.Left) + ")"
	} else {
		r += "()"
	}
	if root.Right != nil {
		r += "(" + tree2str(root.Right) + ")"
	}
	return r
}
