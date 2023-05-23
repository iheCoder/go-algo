package delete_node_in_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	p := root
	var parent *TreeNode
	var isLeft bool
	for p != nil {
		if p.Val == key {
			break
		}
		parent = p
		if p.Val > key {
			isLeft = true
			p = p.Left
		} else {
			isLeft = false
			p = p.Right
		}
	}
	if p == nil {
		return root
	}

	dp := dn(p)
	if parent == nil {
		return dp
	}
	if isLeft {
		parent.Left = dp
	} else {
		parent.Right = dp
	}

	return root
}

func dn(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return nil
	}
	if node.Left == nil {
		return node.Right
	} else if node.Right == nil {
		return node.Left
	}

	p := node.Right
	var parent *TreeNode
	for p.Left != nil {
		parent = p
		p = p.Left
	}
	dp := dn(p)
	if parent != nil {
		parent.Left = dp
	} else {
		node.Right = dp
	}
	p.Left = node.Left
	p.Right = node.Right
	return p
}
