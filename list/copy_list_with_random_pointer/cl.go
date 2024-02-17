package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	vals := make([]int, 0)
	p := head
	var i int
	for p != nil {
		vals = append(vals, p.Val)
		p.Val = i
		i++
		p = p.Next
	}

	m := make(map[int]int)
	p = head
	for p != nil {
		if p.Random != nil {
			m[p.Val] = p.Random.Val
		}
		p = p.Next
	}

	n := len(vals)
	nodes := make([]*Node, n)
	nodes[0] = &Node{Val: vals[0]}
	for j := 1; j < n; j++ {
		nodes[j] = &Node{Val: vals[j]}
		nodes[j-1].Next = nodes[j]
	}

	for k, v := range m {
		nodes[k].Random = nodes[v]
	}

	p = head
	for _, val := range vals {
		p.Val = val
		p = p.Next
	}
	return nodes[0]
}
