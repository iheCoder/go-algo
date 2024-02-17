package all_possible_full_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 || n < 1 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{}}
	}
	if n == 3 {
		return []*TreeNode{{Left: &TreeNode{}, Right: &TreeNode{}}}
	}
	var ans []*TreeNode
	for i := 1; i < n; i += 2 {
		lts := allPossibleFBT(i)
		rts := allPossibleFBT(n - 1 - i)
		for _, lt := range lts {
			for _, rt := range rts {
				root := &TreeNode{}
				root.Left = lt
				root.Right = rt
				ans = append(ans, root)
			}
		}
	}
	return ans
}

func constructTree(n int) *TreeNode {
	if n == 1 {
		return &TreeNode{}
	}
	if n == 3 {
		return &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}
	}
	root := &TreeNode{}
	for i := 1; i < n; i += 2 {
		root.Left = constructTree(i)
		root.Right = constructTree(n - i - 1)
	}
	return root
}
