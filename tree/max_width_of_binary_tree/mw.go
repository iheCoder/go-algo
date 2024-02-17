package max_width_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r := 1
	root.Val = 1
	stk := []*TreeNode{root}
	vc := 1
	var p *TreeNode
	start := 1
	for len(stk) != 0 {
		p = stk[0]
		stk = stk[1:]
		vc--

		if p.Left != nil {
			p.Left.Val = p.Val*2 - 1
			stk = append(stk, p.Left)
		}
		if p.Right != nil {
			p.Right.Val = p.Val * 2
			stk = append(stk, p.Right)
		}

		if vc == 0 {
			r = maxInt(r, p.Val-start+1)
			vc = len(stk)
			if len(stk) > 0 {
				start = stk[0].Val
			}
		}
	}

	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//func widthOfBinaryTree(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	r := 1
//	stk := []*TreeNode{root}
//	vc := 1
//	var p *TreeNode
//	var l, tl int
//	for len(stk) != 0 {
//		p = stk[0]
//		stk = stk[1:]
//		vc--
//
//		if p.Left != nil {
//			l++
//			l += tl
//			tl = 0
//			stk = append(stk, p.Left)
//		} else {
//			tl++
//		}
//		if p.Right != nil {
//			l++
//			l += tl
//			tl = 0
//			stk = append(stk, p.Right)
//		} else {
//			tl++
//		}
//
//		if vc == 0 {
//			r = maxInt(r, l)
//			l = 0
//			tl = 0
//			vc = len(stk)
//		}
//	}
//
//	return r
//}
