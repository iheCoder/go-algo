package sum_of_root_to_leaf_binary_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var r uint32
	var f func(n *TreeNode, sum uint32, d int)
	f = func(n *TreeNode, sum uint32, d int) {
		if n == nil {
			r += reverseBits(sum, d)
			return
		}
		if n.Val == 1 {
			sum += 1 << d
		}
		if n.Left == nil && n.Right == nil {
			r += reverseBits(sum, d)
			return
		}
		if n.Left != nil {
			f(n.Left, sum, d+1)
		}
		if n.Right != nil {
			f(n.Right, sum, d+1)
		}
	}
	f(root, 0, 0)

	return int(r)
}

func reverseBits(x uint32, d int) uint32 {
	if x == 0 {
		return x
	}
	var r uint32
	for i := 0; i < d+1; i++ {
		r |= x & 1 << (d - i)
		x >>= 1
	}
	return r
}

//func sumR2L(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return [][]int{}
//	}
//
//	r := append(sumR2L(root.Left), sumR2L(root.Right)...)
//	for i := 0; i < len(r); i++ {
//		r[i] = append(r[i], root.Val)
//	}
//
//	return r
//}
