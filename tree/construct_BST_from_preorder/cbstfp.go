package construct_BST_from_preorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var border int
	for border = 1; border < len(preorder); border++ {
		if preorder[border] > preorder[0] {
			break
		}
	}

	root.Left = bstFromPreorder(preorder[1:border])
	if border < len(preorder) {
		root.Right = bstFromPreorder(preorder[border:])
	}
	return root
}
