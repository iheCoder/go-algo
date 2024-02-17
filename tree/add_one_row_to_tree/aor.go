package add_one_row_to_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	stk := []*TreeNode{root}
	d := 2
	vc := 1
	var p *TreeNode
	for len(stk) != 0 {
		if d == depth {
			for len(stk) != 0 {
				p = stk[0]
				stk = stk[1:]
				l := &TreeNode{
					Val:  val,
					Left: p.Left,
				}
				p.Left = l

				r := &TreeNode{
					Val:   val,
					Right: p.Right,
				}
				p.Right = r
			}
			return root
		}

		p = stk[0]
		stk = stk[1:]
		vc--
		if p.Left != nil {
			stk = append(stk, p.Left)
		}
		if p.Right != nil {
			stk = append(stk, p.Right)
		}
		if vc == 0 {
			vc = len(stk)
			d++
		}
	}

	return root
}
