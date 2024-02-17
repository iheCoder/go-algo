package populating_next_right_pointers_in_each_node_ii

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	stk := []*Node{root}
	for len(stk) > 0 {
		q := stk
		l := &Node{}
		stk = nil
		for i := 0; i < len(q); i++ {
			l.Next = q[i]
			l = q[i]
			if l.Left != nil {
				stk = append(stk, l.Left)
			}
			if l.Right != nil {
				stk = append(stk, l.Right)
			}
		}
	}
	return root
}

// 	var f func(n *Node) (*Node, *Node)
//	f = func(n *Node) (*Node, *Node) {
//		if n == nil {
//			return nil, nil
//		}
//		if n.Left == nil && n.Right == nil {
//			return n, n
//		}
//
//		if n.Left != nil {
//			n.Left.Next = n.Right
//		}
//
//		lcl, lcr := f(n.Left)
//		rcl, rcr := f(n.Right)
//		left, right := lcl, rcr
//		if left == nil {
//			left = lcr
//			if left == nil {
//				left = rcl
//				if left == nil {
//					left = rcr
//				}
//			}
//		}
//		if right == nil {
//			right = rcl
//			if right == nil {
//				right = lcr
//				if right == nil {
//					right = lcl
//				}
//			}
//		}
//
//		cl, cr := lcl, rcr
//		if cl == nil {
//			cl = lcr
//		}
//		if cr == nil {
//			cr = rcl
//		}
//		if cl != nil {
//			cl.Next = cr
//		}
//
//		// 第一个参数返回如果是左子树的最右节点
//		// 第二个参数返回如果是右子树的最左节点
//		return right, left
//	}
//	f(root)
//	return root
//func connect(root *Node) *Node {
//	var f func(n *Node) (*Node, *Node)
//	f = func(n *Node) (*Node, *Node) {
//		if n == nil {
//			return nil, nil
//		}
//
//		lr, rl := n.Right, n.Right
//		if n.Left != nil {
//			n.Left.Next = n.Right
//			lr = n.Right
//			rl = n.Left
//		}
//		nlr, _ := f(n.Left)
//		_, nrl := f(n.Right)
//		if nlr != nil {
//			nlr.Next = nrl
//		}
//		return lr, rl
//	}
//	f(root)
//	return root
//}
