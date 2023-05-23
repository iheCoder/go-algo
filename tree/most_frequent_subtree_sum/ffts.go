package most_frequent_subtree_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)
	mc := 0
	var f func(n *TreeNode) int
	f = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		sum := n.Val + f(n.Left) + f(n.Right)
		m[sum]++
		mc = maxInt(mc, m[sum])
		return sum
	}
	f(root)

	r := make([]int, 0)
	for k, v := range m {
		if v == mc {
			r = append(r, k)
		}
	}
	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
