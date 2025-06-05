package find_elements_in_a_contaminated_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	m map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	m := make(map[int]bool)
	var reduction func(n *TreeNode, x int)
	reduction = func(n *TreeNode, x int) {
		if n == nil {
			return
		}

		m[x] = true
		lv, rv := 2*x+1, 2*x+2
		reduction(n.Left, lv)
		reduction(n.Right, rv)
	}
	reduction(root, 0)

	return FindElements{m: m}
}

func (fe *FindElements) Find(target int) bool {
	return fe.m[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
