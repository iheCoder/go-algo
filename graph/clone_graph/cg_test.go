package clone_graph

import "testing"

func TestCG(t *testing.T) {
	n := &Node{
		Val: 1,
	}
	nb := []*Node{
		{
			Val:       2,
			Neighbors: []*Node{n},
		},
		{
			Val: 3,
		},
	}
	n.Neighbors = nb

	//cg(n)
	n = &Node{Val: 1}
	nb = []*Node{
		{
			Val:       2,
			Neighbors: []*Node{n},
		},
	}
	n.Neighbors = nb

	cg(n)
}
