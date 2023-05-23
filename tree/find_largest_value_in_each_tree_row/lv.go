package find_largest_value_in_each_tree_row

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	vc := 1
	var p *TreeNode
	var r []int
	var mv *int
	for len(stack) != 0 {
		p = stack[0]
		stack = stack[1:]
		vc--
		mv = maxInt(mv, p.Val)

		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}

		if vc == 0 {
			vc = len(stack)
			r = append(r, *mv)
			mv = nil
		}
	}

	return r
}

func maxInt(mv *int, x int) *int {
	if mv == nil {
		return &x
	}
	if x > *mv {
		return &x
	}
	return mv
}
