package unique_binary_search_trees_ii

func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{
			{
				Val: n,
			},
		}
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	r := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, ln := range left {
			for _, rn := range right {
				root := &TreeNode{
					Val:   i,
					Left:  ln,
					Right: rn,
				}
				r = append(r, root)
			}
		}
	}

	return r
}
