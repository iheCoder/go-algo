package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return sa(nums)
}

func sa(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	low, high := 0, len(nums)-1
	mid := (high-low)>>1 + low
	root := &TreeNode{
		Val: nums[mid],
	}

	if mid > 0 {
		root.Left = sa(nums[:mid])
	}
	if mid+1 < len(nums) {
		root.Right = sa(nums[mid+1:])
	}

	return root
}
