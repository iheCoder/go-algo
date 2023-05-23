package construct_binary_tree_from_iorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rv := postorder[len(postorder)-1]
	root := &TreeNode{Val: rv}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == rv {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
	root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
	return root
}
