package max_binary_tree

func constructMaximumBinaryTreeMS(nums []int) *TreeNode {
	n := len(nums)
	// left、right记录元素在nums中左右边界
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i] = -1
	}
	tree := make([]*TreeNode, n)

	stk := []int{-1}
	for i, num := range nums {
		tree[i] = &TreeNode{Val: num}
		for len(stk) > 1 && num > nums[stk[len(stk)-1]] {
			right[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
		left[i] = stk[len(stk)-1]
		stk = append(stk, i)
	}

	var root *TreeNode
	for i, node := range tree {
		l, r := left[i], right[i]
		if l == -1 && r == -1 {
			root = node
		} else if r == -1 || l != -1 && nums[l] < nums[r] {
			tree[l].Right = node
		} else {
			tree[r].Left = node
		}
	}
	return root
}

// func constructMaximumBinaryTree(nums []int) *TreeNode {
//    tree := make([]*TreeNode, len(nums))
//    stk := []int{}
//    for i, num := range nums {
//        tree[i] = &TreeNode{Val: num}
//        for len(stk) > 0 && num > nums[stk[len(stk)-1]] {
//            tree[i].Left = tree[stk[len(stk)-1]]
//            stk = stk[:len(stk)-1]
//        }
//        if len(stk) > 0 {
//            tree[stk[len(stk)-1]].Right = tree[i]
//        }
//        stk = append(stk, i)
//    }
//    return tree[stk[0]]
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/maximum-binary-tree/solution/zui-da-er-cha-shu-by-leetcode-solution-lbeo/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
