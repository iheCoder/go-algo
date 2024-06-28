package complete_binary_tree_inserter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root                 *TreeNode
	prevLevel, nextLevel []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	prevLevel := make([]*TreeNode, 0)
	nextLevel := []*TreeNode{root}

	for len(prevLevel) == 0 || prevLevel[len(prevLevel)-1].Right != nil {
		prevLevel = nextLevel
		nextLevel = nil
		for _, node := range prevLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
	}

	return CBTInserter{
		root:      root,
		prevLevel: prevLevel,
		nextLevel: nextLevel,
	}
}

func (i *CBTInserter) Insert(val int) int {
	newNode := &TreeNode{
		Val: val,
	}

	if len(i.prevLevel) == 0 {
		i.prevLevel = i.nextLevel
		i.nextLevel = make([]*TreeNode, 0)
	}

	i.nextLevel = append(i.nextLevel, newNode)

	node := i.prevLevel[0]
	j := 1
	for node.Left != nil && node.Right != nil {
		node = i.prevLevel[j]
		j++
	}

	if node.Left == nil {
		node.Left = newNode
		return node.Val
	}

	node.Right = newNode
	i.prevLevel = i.prevLevel[j:]
	return node.Val
}

func (i *CBTInserter) Get_root() *TreeNode {
	return i.root
}
