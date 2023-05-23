package diameter_of_binary_tree

import (
	"testing"
)

func TestDBT(t *testing.T) {
	root := &TreeNode{Val: 1}
	t.Log(diameterOfBinaryTree(root))
}
