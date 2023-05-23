package two_sum_iv_input_is_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 显然并不能像在数组那样用上双指针了
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	m := make(map[int]int)
	q := make([]*TreeNode, 1)
	q[0] = root
	var p *TreeNode

	for len(q) != 0 {
		p = q[0]
		q = q[1:]
		m[p.Val]++
		if p.Left != nil {
			q = append(q, p.Left)
		}
		if p.Right != nil {
			q = append(q, p.Right)
		}
	}

	for v, c := range m {
		if 2*v == k {
			if c > 1 {
				return true
			}
			continue
		}
		if _, ok := m[k-v]; ok {
			return true
		}
	}

	return false
}
