package flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	var f func(n *Node) (*Node, *Node)
	f = func(n *Node) (*Node, *Node) {
		p := n
		var next, prev *Node
		for p != nil {
			next = p.Next
			if p.Child != nil {
				head, tail := f(p.Child)
				p.Next = head
				head.Prev = p
				tail.Next = next
				if next != nil {
					next.Prev = tail
				}
				p.Child = nil
				prev = tail
			} else {
				prev = p
			}
			// 由于只引用了本层，所以导致tail返回的就是原本层的最后一个
			//prev = p
			p = next
		}

		return n, prev
	}
	f(root)
	return root
}

// 原来是规定了顺序的啊！
//func flatten(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//
//	p := root
//	var child, head, prev *Node
//	for p != nil {
//		child = nil
//		for p != nil {
//			if p.Child != nil {
//				if child == nil {
//					child = p.Child
//				}
//				p.Child = nil
//			}
//			prev = p
//			p = p.Next
//		}
//
//		if child == nil {
//			break
//		}
//
//		head = child
//		for head.Prev != nil {
//			head = head.Prev
//		}
//		prev.Next = head
//		head.Prev = prev
//		p = head
//	}
//
//	return root
//}
