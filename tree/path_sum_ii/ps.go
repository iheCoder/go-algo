package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	return ps(root, targetSum, 0, make([]int, 0))
}

func ps(root *TreeNode, targetSum, sum int, arr []int) [][]int {
	if root == nil {
		return [][]int{}
	}

	narr := make([]int, len(arr))
	copy(narr, arr)
	narr = append(narr, root.Val)
	sum += root.Val

	if root.Left == nil && root.Right == nil {
		if sum == targetSum {
			return [][]int{narr}
		}
		return [][]int{}
	}

	l := ps(root.Left, targetSum, sum, narr)
	r := ps(root.Right, targetSum, sum, narr)

	return append(l, r...)
}
