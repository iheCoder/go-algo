package average_of_levels_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (averages []float64) {
	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		sum := 0
		curLevel := nextLevel
		nextLevel = nil
		for _, node := range curLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(len(curLevel)))
	}
	return
}
