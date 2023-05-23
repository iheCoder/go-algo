package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ColorfulTreeNode struct {
	c    color
	node *TreeNode
}

type color int

const (
	white color = iota
	gray
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	r := make([]int, 0)
	stack := make([]*ColorfulTreeNode, 1)
	stack[0] = &ColorfulTreeNode{
		c:    white,
		node: root,
	}
	var p *ColorfulTreeNode
	var left *ColorfulTreeNode
	for len(stack) != 0 {
		p = stack[len(stack)-1]

		if p.c == white && p.node.Left != nil {
			p.c = gray
			for p.node.Left != nil {
				left = &ColorfulTreeNode{
					c:    gray,
					node: p.node.Left,
				}
				stack = append(stack, left)
				p = left
			}
		}

		r = append(r, p.node.Val)
		stack = stack[:len(stack)-1]

		if p.node.Right != nil {
			stack = append(stack, &ColorfulTreeNode{
				c:    white,
				node: p.node.Right,
			})
		}

	}

	return r
}

// 		if p != nil {
//			stack = append(stack, p)
//			left = p.Left
//			for left != nil {
//				stack = append(stack, left)
//				p = left
//				left = left.Left
//			}
//		}
