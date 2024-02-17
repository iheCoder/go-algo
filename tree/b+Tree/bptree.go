package b_Tree

//type Node struct {
//	order  int64
//	values []int64
//	nodes  []*Node
//	next   *Node
//	parent *Node
//	isLeaf bool
//}
//
//func (n *Node) InsertLeaf(value int64, node *Node) {
//	if len(n.values) != 0 {
//		for i, v := range n.values {
//			// key相等时，仅仅只是添加一个索引
//			if value == v {
//				n.nodes = append(n.nodes, node)
//				break
//			} else if value < v {
//				n.nodes =append(n.nodes[:i+1],n.nodes[i:]...)
//				n.nodes[i]= node
//				n.values=append(n.values[:i+1],n.values[i:]...)
//				n.values[i]=value
//				break
//			}else if i+1 == len(n.values){
//				n.nodes = append(n.nodes, node)
//				n.values = append(n.values, value)
//				break
//			}
//		}
//	} else {
//		n.values = make([]int64, 1)
//		n.nodes = make([]*Node, 1)
//		n.nodes[0] = node
//		n.values[0] = value
//	}
//}
//
//type BPlusTree struct {
//	root *Node
//	order int
//}
//
//func (t *BPlusTree) find(value int64, n *Node) bool {
//
//}
//
//func (t *BPlusTree) search(value int64) *Node {
//	n:=t.root
//	for !n.isLeaf {
//		for i, v := range n.values {
//			if value == v {
//				n=n.nodes[i]
//				break
//			}else if value < v{
//				n=n.nodes[i-1]
//				break
//			}else if i == len(n.values)-1 {
//				n=n.nodes[i+1]
//			}
//		}
//	}
//	return n
//}
//
//func (t *BPlusTree) insert(key, value int64) {
//
//}
//
//func NewBPlusTree(order int) *BPlusTree {
//	return &BPlusTree{
//		root:  &Node{},
//		order: order,
//	}
//}
