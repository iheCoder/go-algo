package min_cost_tree_from_leaf_values

func mctFromLeafValues(arr []int) int {
	var stk []int
	res := 0
	for _, v := range arr {
		for len(stk) > 0 && stk[len(stk)-1] < v {
			if len(stk) == 1 || stk[len(stk)-2] > v {
				res += stk[len(stk)-1] * v
			} else {
				res += stk[len(stk)-1] * stk[len(stk)-2]
			}
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, v)
	}

	for len(stk) > 1 {
		res += stk[len(stk)-1] * stk[len(stk)-2]
		stk = stk[:len(stk)-1]
	}

	return res
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// 不一定非完全二叉树吧
// 	n := len(arr)
//	var sub int
//	if n%2 == 1 {
//		arr = append(arr, 1)
//	}
//
//	sort.Ints(arr)
//	if n%2 == 1 {
//		sub = arr[len(arr)-1]
//	}
//	tmpArr := make([]int, 0, len(arr)/2)
//	var ans int
//	match := func() bool {
//		if len(arr) == 1 {
//			return true
//		}
//
//		i, j := 0, len(arr)-1
//		var p int
//		for i < j {
//			p = arr[i] * arr[j]
//			ans += p
//			tmpArr = append(tmpArr, maxInt(arr[i], arr[j]))
//			i++
//			j--
//		}
//
//		arr = tmpArr
//		sort.Ints(arr)
//		tmpArr = make([]int, 0, len(arr)/2)
//		return false
//	}
//
//	for {
//		if match() {
//			return ans - sub
//		}
//	}
