package leaf_similar_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	}
	stk1 := make([]*TreeNode, 0)
	stk2 := make([]*TreeNode, 0)
	p1 := root1
	p2 := root2
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for p1 != nil || len(stk1) != 0 {
		for p1 != nil {
			stk1 = append(stk1, p1)
			p1 = p1.Left
		}
		p1 = stk1[len(stk1)-1]
		stk1 = stk1[:len(stk1)-1]
		if p1.Left == nil && p1.Right == nil {
			arr1 = append(arr1, p1.Val)
		}
		p1 = p1.Right
	}
	for p2 != nil || len(stk2) != 0 {
		for p2 != nil {
			stk2 = append(stk2, p2)
			p2 = p2.Left
		}
		p2 = stk2[len(stk2)-1]
		stk2 = stk2[:len(stk2)-1]
		if p2.Left == nil && p2.Right == nil {
			arr2 = append(arr2, p2.Val)
		}
		p2 = p2.Right
	}

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
