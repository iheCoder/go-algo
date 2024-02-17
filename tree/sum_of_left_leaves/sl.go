package sum_of_left_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}

	return sll(root, 0)
}

func sll(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}

	left := node.Left
	if left != nil {
		if left.Left == nil && left.Right == nil {
			sum += left.Val
		} else {
			sum = sll(left, sum)
		}
	}

	return sll(node.Right, sum)
}
