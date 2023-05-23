package find_dup_subtrees

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	repeat := make(map[*TreeNode]bool)
	seen := make(map[string]*TreeNode)
	var dfs func(n *TreeNode) string
	dfs = func(n *TreeNode) string {
		if n == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", n.Val, dfs(n.Left), dfs(n.Right))
		if rn, ok := seen[serial]; ok {
			repeat[rn] = true
		} else {
			seen[serial] = n
		}
		return serial
	}
	dfs(root)

	r := make([]*TreeNode, 0, len(repeat))
	for node, _ := range repeat {
		r = append(r, node)
	}
	return r
}
