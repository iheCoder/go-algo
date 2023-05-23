package sum_root_to_leaf_num

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return sn(root, "", 0)
}

func sn(node *TreeNode, prefix string, totalSum int) int {
	if node == nil {
		return totalSum
	}

	prefix += strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		sum, _ := strconv.Atoi(prefix)
		return sum + totalSum
	}

	totalSum = sn(node.Left, prefix, totalSum)
	totalSum = sn(node.Right, prefix, totalSum)
	return totalSum
}
