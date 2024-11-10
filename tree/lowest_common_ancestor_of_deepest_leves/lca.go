package lowest_common_ancestor_of_deepest_leves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	cr, _ := ldl(root, 0)
	return cr
}

func ldl(node *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return nil, depth
	}

	depth++

	if node.Left == nil && node.Right == nil {
		return node, depth
	}

	var ld, rd int
	var ln, rn *TreeNode
	if node.Left != nil {
		ln, ld = ldl(node.Left, depth)
	}
	if node.Right != nil {
		rn, rd = ldl(node.Right, depth)
	}

	if ld > rd {
		return ln, ld
	} else if rd > ld {
		return rn, rd
	}

	return node, ld
}
