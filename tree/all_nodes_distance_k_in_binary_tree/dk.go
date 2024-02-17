package all_nodes_distance_k_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if target == nil {
		return []int{}
	}
	r, _ := dk(root, target, k, -1)
	return r
}

func dk(root, target *TreeNode, k, rk int) ([]int, int) {
	if root == nil {
		return []int{}, -1
	}

	var r []int
	var lr, rr []int
	var ld, rd int
	if rk != -1 {
		rk++
		if rk == k {
			r = append(r, root.Val)
		} else if rk < k {
			lr, _ = dk(root.Left, target, k, rk)
			rr, _ = dk(root.Right, target, k, rk)
			r = append(r, lr...)
			r = append(r, rr...)
		}
	} else {
		if root == target {
			rk = 0
			lr, _ = dk(root.Left, target, k, rk)
			rr, _ = dk(root.Right, target, k, rk)
		} else {
			lr, ld = dk(root.Left, target, k, rk)
			rr, rd = dk(root.Right, target, k, rk)
			if ld == -1 && rd == -1 {
				return []int{}, -1
			} else if ld != -1 {
				rk = ld + 1
				rr, rd = dk(root.Right, target, k, rk)
			} else if rd != -1 {
				rk = rd + 1
				lr, ld = dk(root.Left, target, k, rk)
			}
		}
		if rk == k {
			r = append(r, root.Val)
		}
		r = append(r, lr...)
		r = append(r, rr...)
	}
	return r, rk
}
