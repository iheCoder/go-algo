package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	ft(root)
}

func ft(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	tail := ft(root.Right)
	if root.Left != nil {
		leftTail := ft(root.Left)
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		if tail == nil {
			tail = leftTail
		}
	}
	return tail
}
