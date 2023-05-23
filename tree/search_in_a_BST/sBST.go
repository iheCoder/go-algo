package search_in_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	p := root
	for p != nil {
		if val == p.Val {
			return p
		}
		if p.Val > val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}
