package balance_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return build(inorder(root))
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	l := inorder(root.Left)
	r := inorder(root.Right)
	l = append(l, root.Val)
	return append(l, r...)
}

func build(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	root := &TreeNode{Val: arr[mid]}
	if mid > 0 {
		root.Left = build(arr[:mid])
	}
	if mid+1 < len(arr) {
		root.Right = build(arr[mid+1:])
	}
	return root
}
