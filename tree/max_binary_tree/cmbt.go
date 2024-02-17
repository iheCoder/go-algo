package max_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	i := maxAt(nums)
	root := &TreeNode{Val: nums[i]}
	root.Left = constructMaximumBinaryTree(nums[:i])
	root.Right = constructMaximumBinaryTree(nums[i+1:])
	return root
}

func maxAt(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}
