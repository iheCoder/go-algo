package convert_bst_to_greater_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var suffixSum int
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Right
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		suffixSum += p.Val
		p.Val = suffixSum
		p = p.Left
	}

	return root
}
