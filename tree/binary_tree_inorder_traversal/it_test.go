package binary_tree_inorder_traversal

import "testing"

const NullNode = 1<<31 - 1

func TestBuildBinaryTree(t *testing.T) {
	arr := []int{1, NullNode, 2, 3}
	buildBinaryTree(arr)
}

func TestIT(t *testing.T) {
	arr := []int{1, 2, NullNode, 3, 4}
	root := buildBinaryTree(arr)
	r := inorderTraversal(root)
	t.Log(r)
}

func buildBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	var left, right *TreeNode
	var p *TreeNode
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for i := 1; i < len(arr); i += 2 {
		p = stack[0]
		stack = stack[1:]
		if arr[i] != NullNode {
			left = &TreeNode{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}
			p.Left = left
			stack = append(stack, left)
		}
		if i+1 < len(arr) && arr[i+1] != NullNode {
			right = &TreeNode{
				Val:   arr[i+1],
				Left:  nil,
				Right: nil,
			}
			p.Right = right
			stack = append(stack, right)
		}
	}

	return root
}
