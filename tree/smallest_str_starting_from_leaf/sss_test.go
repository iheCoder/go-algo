package smallest_str_starting_from_leaf

import "testing"

func TestSSS(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  0,
			Left: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 1},
	}
	t.Log(smallestFromLeaf(root))
}
