package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		}
		return false
	}

	var a, b *TreeNode
	as := make([]*TreeNode, 1)
	as[0] = p
	bs := make([]*TreeNode, 1)
	bs[0] = q
	for len(as) != 0 && len(bs) != 0 {
		a = as[0]
		as = as[1:]
		b = bs[0]
		bs = bs[1:]
		if a.Val != b.Val {
			return false
		}

		if a.Left != nil && b.Left != nil {
			as = append(as, a.Left)
			bs = append(bs, b.Left)
		} else if !(a.Left == nil && b.Left == nil) {
			return false
		}
		if a.Right != nil && b.Right != nil {
			as = append(as, a.Right)
			bs = append(bs, b.Right)
		} else if !(a.Right == nil && b.Right == nil) {
			return false
		}
	}

	return len(as) == 0 && len(bs) == 0
}
