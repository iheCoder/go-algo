package path_sum_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return ps(root, targetSum, make([]int, 0))
}

func ps(root *TreeNode, targetSum int, prefix []int) int {
	if root == nil {
		return 0
	}

	count := 0
	if root.Val == targetSum {
		count++
	}
	if len(prefix) != 0 {
		for i := 0; i < len(prefix); i++ {
			prefix[i] += root.Val
			if prefix[i] == targetSum {
				count++
			}
		}
	}
	prefix = append(prefix, root.Val)
	pl := make([]int, len(prefix))
	copy(pl, prefix)
	pr := make([]int, len(prefix))
	copy(pr, prefix)

	lc := ps(root.Left, targetSum, pl)
	rc := ps(root.Right, targetSum, pr)
	return count + lc + rc
}
