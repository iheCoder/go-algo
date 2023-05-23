package trim_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val <= high && root.Val >= low {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	} else if root.Val > high {
		return trimBST(root.Left, low, high)
	} else {
		return trimBST(root.Right, low, high)
	}
}
