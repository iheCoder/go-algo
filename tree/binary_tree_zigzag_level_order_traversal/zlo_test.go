package binary_tree_zigzag_level_order_traversal

import "testing"

const NullNode = 1<<31 - 1

func TestZLO(t *testing.T) {
	arr := []int{0, 2, 4, 1, NullNode, 3, -1, 5, 1, NullNode, 6, NullNode, 8}
	root := buildBinaryTree(arr)
	t.Log(zigzagLevelOrder(root))
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
