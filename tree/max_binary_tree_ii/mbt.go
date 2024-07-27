package max_binary_tree_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{
					Val:  val,
					Left: root,
				}
			}

			parent.Right = &TreeNode{
				Val:  val,
				Left: cur,
			}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{
		Val: val,
	}
	return root
}
