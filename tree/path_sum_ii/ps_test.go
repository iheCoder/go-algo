package path_sum_ii

import "testing"

func TestPS(t *testing.T) {
	arr := []int{5, 4, 8, 11, NullNode, 13, 4, 7, 2, NullNode, NullNode, 5, 1}
	root := buildBinaryTree(arr)
	pathSum(root, 22)
}

const NullNode = 1<<31 - 1

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
