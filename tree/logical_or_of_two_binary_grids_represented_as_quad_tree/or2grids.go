package logical_or_of_two_binary_grids_represented_as_quad_tree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	var f func(t1, t2 *Node) *Node
	f = func(t1, t2 *Node) *Node {
		if t1.IsLeaf && t2.IsLeaf {
			return &Node{
				Val:    t1.Val || t2.Val,
				IsLeaf: true,
			}
		}

		n := &Node{}
		tl1, tl2 := t1.TopLeft, t2.TopLeft
		if tl1 == nil {
			tl1 = t1
		}
		if tl2 == nil {
			tl2 = t2
		}
		n.TopLeft = f(tl1, tl2)

		tr1, tr2 := t1.TopRight, t2.TopRight
		if tr1 == nil {
			tr1 = t1
		}
		if tr2 == nil {
			tr2 = t2
		}
		n.TopRight = f(tr1, tr2)

		bl1, bl2 := t1.BottomLeft, t2.BottomLeft
		if bl1 == nil {
			bl1 = t1
		}
		if bl2 == nil {
			bl2 = t2
		}
		n.BottomLeft = f(bl1, bl2)

		br1, br2 := t1.BottomRight, t2.BottomRight
		if br1 == nil {
			br1 = t1
		}
		if br2 == nil {
			br2 = t2
		}
		n.BottomRight = f(br1, br2)

		if n.TopLeft.IsLeaf && n.TopRight.IsLeaf && n.BottomLeft.IsLeaf && n.BottomRight.IsLeaf {
			if n.TopLeft.Val == n.TopRight.Val && n.BottomLeft.Val == n.BottomRight.Val && n.TopLeft.Val == n.BottomLeft.Val {
				return &Node{Val: n.TopLeft.Val, IsLeaf: true}
			}
		}

		return n
	}
	return f(quadTree1, quadTree2)
}
