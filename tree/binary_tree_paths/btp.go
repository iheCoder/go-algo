package binary_tree_paths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	prefix := strconv.Itoa(root.Val)

	lps := btp(root.Left, prefix)
	rps := btp(root.Right, prefix)
	return append(lps, rps...)
}

func btp(node *TreeNode, prefix string) []string {
	if node == nil {
		return nil
	}

	prefix += "->" + strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		return []string{prefix}
	}

	lps := btp(node.Left, prefix)
	rps := btp(node.Right, prefix)
	return append(lps, rps...)
}
