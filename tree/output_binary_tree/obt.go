package output_binary_tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	stk := []*TreeNode{root}
	var p *TreeNode
	vc := 1
	var h int
	for len(stk) != 0 {
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
			h++
		}
	}

	col := (1 << h) - 1
	r := make([][]string, h)
	for i := 0; i < len(r); i++ {
		r[i] = make([]string, col)
	}
	fillMatrix(r, 0, col-1, 0, root)
	return r
}

func fillMatrix(m [][]string, i, j, k int, root *TreeNode) {
	if root == nil {
		return
	}

	mid := i + (j-i)>>1
	m[k][mid] = strconv.Itoa(root.Val)
	if root.Left != nil {
		fillMatrix(m, i, mid-1, k+1, root.Left)
	}
	if root.Right != nil {
		fillMatrix(m, mid+1, j, k+1, root.Right)
	}
}

//func pt(root *TreeNode) [][]string {
//	if root == nil {
//		return nil
//	}
//
//	r := append(printTree(root.Left), printTree(root.Right)...)
//	rs := make([]string, len(r))
//	mid := len(rs) / 2
//	rs[mid]= strconv.Itoa(root.Val)
//	r = append(r, )
//}
