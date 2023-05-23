package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// (1 get level order, then reverse
// (2 travel two times at same time
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	stack := []*TreeNode{root}
	r := make([][]int, 1)
	var p *TreeNode
	levelCnt := 1
	var level int
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
			levelCnt = len(stack)
			level++
			r = append(r, make([]int, 0, levelCnt))
		}
	}

	var low, high int
	for i := 1; i < len(r); i += 2 {
		low, high = 0, len(r[i])-1
		for low < high {
			r[i][low], r[i][high] = r[i][high], r[i][low]
			low++
			high--
		}
	}

	return r
}

//func zigzagLevelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return [][]int{}
//	}
//
//	stack := []*TreeNode{root}
//	r := make([][]int, 0)
//	var level int
//	var p *TreeNode
//	count := 1
//	for count != 0 {
//		r = append(r, make([]int, 0, count))
//		for i := count - 1; i >= 0; i-- {
//			p = stack[i]
//			r[level] = append(r[level], p.Val)
//			if p.Right != nil {
//				stack = append(stack, p.Right)
//			}
//			if p.Left != nil {
//				stack = append(stack, p.Left)
//			}
//		}
//
//		stack = stack[count:]
//		count = len(stack)
//		if count <= 0 {
//			return r
//		}
//		r = append(r, make([]int, 0, count))
//		level++
//		for i := 0; i < count; i++ {
//			p = stack[i]
//			r[level] = append(r[level], p.Val)
//			rp := stack[count-1-i]
//			if rp.Right != nil {
//				stack = append(stack, rp.Right)
//			}
//			if rp.Left != nil {
//				stack = append(stack, rp.Left)
//			}
//		}
//
//		stack = stack[count:]
//		count = len(stack)
//		level++
//	}
//
//	return r
//}

// 	for len(stack) != 0 {
//
//		for _, p := range stack {
//			r[level] = append(r[level], p.Val)
//
//		}
//		p = stack[0]
//		stack = stack[1:]
//		if p == nil {
//			continue
//		}
//		r[level] = append(r[level], p.Val)
//
//		if p.Right != nil {
//			r[level] = append(r[level], p.Right.Val)
//			stack = append(stack, p.Right.Right)
//			stack = append(stack, p.Right.Left)
//		}
//		if p.Left != nil {
//			r[level] = append(r[level], p.Left.Val)
//			stack = append(stack, p.Left.Right)
//			stack = append(stack, p.Left.Left)
//		}
//
//		if verCnt == 0 {
//			level++
//			verCnt = len(stack)
//			r = append(r, make([]int, 0, verCnt))
//		}
//	}
//
//	return r
//}
