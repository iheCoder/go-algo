package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[int]*Node

func cloneGraph(node *Node) *Node {
	visited = make(map[int]*Node)
	return cg(node)
}

func cg(node *Node) *Node {
	if node == nil {
		return nil
	}
	cn := &Node{
		Val: node.Val,
	}
	visited[node.Val] = cn
	nb := make([]*Node, len(node.Neighbors))
	for i, neighbor := range node.Neighbors {
		if ne, ok := visited[neighbor.Val]; ok {
			nb[i] = ne
			continue
		}
		nb[i] = cg(neighbor)
	}
	cn.Neighbors = nb
	return cn
}
