package find_mode_in_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var r []int
	maxLen := 1
	stack := make([]*TreeNode, 0)
	p := root
	var temp, tempCount int
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp != p.Val {
			if tempCount == maxLen {
				r = append(r, temp)
			} else if tempCount > maxLen {
				r = []int{temp}
				maxLen = tempCount
			}
			temp = p.Val
			tempCount = 1
		} else {
			tempCount++
		}
		p = p.Right
	}
	if tempCount == maxLen {
		r = append(r, temp)
	} else if tempCount > maxLen {
		r = []int{temp}
		maxLen = tempCount
	}

	return r
}
