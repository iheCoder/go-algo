package recover_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 一种找到乱序的方法就是——一个人认为我错了，就是我错了。世界认为我错了，便是世界错了
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	nodes := preorder(root)
	var i int
	for i = 0; i < len(nodes)-1 && nodes[i+1].Val > nodes[i].Val; i++ {
	}

	j := len(nodes) - 1
	for ; j > i && nodes[j].Val > nodes[j-1].Val; j-- {
	}

	nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
}

func preorder(root *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 0)
	p := root
	var r []*TreeNode
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, p)
		p = p.Right
	}

	return r
}
