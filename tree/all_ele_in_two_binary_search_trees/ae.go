package all_ele_in_two_binary_search_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	var dfs func(node *TreeNode)
	res := make([]int, 0)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	nums1, nums2 := inorder(root1), inorder(root2)
	n1, n2 := len(nums1), len(nums2)
	merged := make([]int, 0, n1+n2)
	var i, j int

	for {
		if i >= n1 {
			for j < n2 {
				merged = append(merged, nums2[j])
				j++
			}
			break
		} else if j >= n2 {
			for i < n1 {
				merged = append(merged, nums1[i])
				i++
			}
			break
		}

		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	return merged
}
