package subtree_of_another_tree

func isSubtree(s, t *TreeNode) bool {

}

func kmp(s, t []int) bool {
	sl, tl := len(s), len(t)
	fail := make([]int, sl)
	for i := 0; i < sl; i++ {
		fail[i] = -1
	}

	for i, j := 1, -1; i < tl; i++ {
		for j != -1 && t[i] != t[j+1] {
			j = fail[j]
		}
		if t[i] == t[j+1] {
			j++
		}
		fail[i] = j
	}

	for i, j := 0, -1; i < sl; i++ {
		for j != -1 && s[i] != t[j+1] {
			j = fail[j]
		}
		if s[i] == t[j+1] {
			j++
		}
		if j == tl-1 {
			return true
		}
	}
	return false
}

// 获取树t的前序遍历
func getDFSOrder(t *TreeNode, orders []int, lNil, rNil int) []int {
	if t == nil {
		return orders
	}

	orders = append(orders, t.Val)

	if t.Left != nil {
		orders = getDFSOrder(t.Left, orders, lNil, rNil)
	} else {
		orders = append(orders, lNil)
	}
	if t.Right != nil {
		orders = getDFSOrder(t.Right, orders, lNil, rNil)
	} else {
		orders = append(orders, rNil)
	}

	return orders
}

// 获取树t中最大值
func getMaxElement(t *TreeNode, maxEle *int) {
	if t == nil {
		return
	}

	if t.Val > *maxEle {
		*maxEle = t.Val
	}
	getMaxElement(t.Left, maxEle)
	getMaxElement(t.Right, maxEle)
}
