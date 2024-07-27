package smallest_str_starting_from_leaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	var allPath [][]byte
	var f func(n *TreeNode, trace []byte)
	f = func(n *TreeNode, trace []byte) {
		if n == nil {
			allPath = append(allPath, trace)
			return
		}

		trace = append(trace, 'a'+byte(n.Val))
		if n.Left == nil && n.Right == nil {
			allPath = append(allPath, trace)
			return
		} else if n.Left != nil && n.Right == nil {
			f(n.Left, trace)
			return
		} else if n.Left == nil && n.Right != nil {
			f(n.Right, trace)
			return
		}
		ct := make([]byte, len(trace))
		copy(ct, trace)

		f(n.Left, trace)
		f(n.Right, ct)
	}
	f(root, make([]byte, 0))

	if len(allPath) == 0 {
		return ""
	}

	ans := allPath[0]
	for i := 1; i < len(allPath); i++ {
		if len(allPath[i]) == 1 {
			continue
		}
		if len(ans) == 1 {
			ans = allPath[i]
		}
		if less(allPath[i], ans) {
			ans = allPath[i]
		}
	}

	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}

func less(s1, s2 []byte) bool {
	i, j := len(s1)-1, len(s2)-1
	for {
		if i < 0 {
			return true
		}
		if j < 0 {
			return false
		}
		if s1[i] < s2[j] {
			return true
		} else if s1[i] > s2[j] {
			return false
		}
		i--
		j--
	}
}
