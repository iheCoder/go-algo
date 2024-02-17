package construct_binary_tree_from_preorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	lroot := preorder[1]
	var i int
	for ; i < len(postorder); i++ {
		if postorder[i] == lroot {
			break
		}
	}
	lb := i + 2
	if lb > len(preorder) {
		lb = len(preorder)
	}
	root.Left = constructFromPrePost(preorder[1:lb], postorder[:i+1])
	if i+2 < len(preorder) {
		root.Right = constructFromPrePost(preorder[i+2:], postorder[i+1:len(postorder)-1])
	}
	return root
}
