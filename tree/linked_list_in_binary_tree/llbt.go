package linked_list_in_binary_tree

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	allPath := make([]string, 0, 2500)
	var dfs func(n *TreeNode, p string)
	dfs = func(n *TreeNode, p string) {
		if n == nil {
			allPath = append(allPath, p)
			return
		}

		cp := p + strconv.Itoa(n.Val)
		dfs(n.Left, cp)
		dfs(n.Right, cp)
	}
	dfs(root, "")

	target := ""
	p := head
	for p != nil {
		target += strconv.Itoa(p.Val)
		p = p.Next
	}

	for _, s := range allPath {
		if strings.Contains(s, target) {
			return true
		}
	}

	return false
}
