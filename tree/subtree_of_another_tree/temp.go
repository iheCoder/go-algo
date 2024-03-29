package subtree_of_another_tree

import "math"

func isSubtreeKmp(s *TreeNode, t *TreeNode) bool {
	maxEle := math.MinInt32
	getMaxElement(s, &maxEle)
	getMaxElement(t, &maxEle)
	lNull := maxEle + 1
	rNull := maxEle + 2

	sl, tl := getDfsOrder(s, []int{}, lNull, rNull), getDfsOrder(t, []int{}, lNull, rNull)
	return kmp(sl, tl)
}

//func kmp(s, t []int) bool {
//	sLen, tLen := len(s), len(t)
//	fail := make([]int, sLen)
//	for i := 0; i < sLen; i++ {
//		fail[i] = -1
//	}
//	for i, j := 1, -1; i < tLen; i++ {
//		for j != -1 && t[i] != t[j+1] {
//			j = fail[j]
//		}
//		if t[i] == t[j+1] {
//			j++
//		}
//		fail[i] = j
//	}
//
//	for i, j := 0, -1; i < sLen; i++ {
//		for j != -1 && s[i] != t[j+1] {
//			j = fail[j]
//		}
//		if s[i] == t[j+1] {
//			j++
//		}
//		if j == tLen-1 {
//			return true
//		}
//	}
//	return false
//}

func getDfsOrder(t *TreeNode, list []int, lNull, rNull int) []int {
	if t == nil {
		return list
	}
	list = append(list, t.Val)
	if t.Left != nil {
		list = getDfsOrder(t.Left, list, lNull, rNull)
	} else {
		list = append(list, lNull)
	}

	if t.Right != nil {
		list = getDfsOrder(t.Right, list, lNull, rNull)
	} else {
		list = append(list, rNull)
	}
	return list
}
