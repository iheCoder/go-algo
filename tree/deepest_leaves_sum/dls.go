package deepest_leaves_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stk := []*TreeNode{root}
	var p *TreeNode
	vc := 1
	var sum int
	for len(stk) != 0 {
		p = stk[0]
		stk = stk[1:]
		vc--
		sum += p.Val

		if p.Left != nil {
			stk = append(stk, p.Left)
		}
		if p.Right != nil {
			stk = append(stk, p.Right)
		}

		if vc == 0 && len(stk) != 0 {
			vc = len(stk)
			sum = 0
		}
	}
	return sum
}
