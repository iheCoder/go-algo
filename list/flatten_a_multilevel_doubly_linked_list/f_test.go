package flatten_a_multilevel_doubly_linked_list

import "testing"

func TestFlat(t *testing.T) {
	//root := &Node{Val: 1}
	//n2 := &Node{Val: 2, Prev: root}
	//root.Next = n2
	//n3 := &Node{Val: 3}
	//n2.Child = n3
	//flatten(root)

	n3 := &Node{Val: 3}
	n2 := &Node{Val: 2, Child: n3}
	root := &Node{Val: 1, Child: n2}
	flatten(root)
}
