package populating_next_right_pointers_in_each_node_ii

import "testing"

func TestPopulating(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Right: &Node{Val: 7},
		},
	}
	connect(root)
	t.Log("PASS")
}
