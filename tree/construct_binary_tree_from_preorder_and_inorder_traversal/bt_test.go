package construct_binary_tree_from_preorder_and_inorder_traversal

import "testing"

func TestBT(t *testing.T) {
	inorder := []int{10, 9, 3, 15, 20, 7}
	preorder := []int{3, 9, 10, 20, 15, 7}
	buildTree(preorder, inorder)
}
