package binary_tree_level_order_traversal_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	stack := []*TreeNode{root}
	var p *TreeNode
	r := make([][]int, 0)
	r = append(r, make([]int, 0))
	var level int
	levelCnt := 1
	for len(stack) != 0 {
		p = stack[0]
		stack = stack[1:]
		levelCnt--

		r[level] = append(r[level], p.Val)
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}

		if levelCnt == 0 && len(stack) != 0 {
			level++
			levelCnt = len(stack)
			r = append(r, make([]int, 0))
		}
	}

	i, j := 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return r
}
