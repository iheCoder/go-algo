package add_two_negabinary_num

func addNegabinary(arr1 []int, arr2 []int) []int {
	i := len(arr1) - 1
	j := len(arr2) - 1
	var carry int
	var ans []int
	for i >= 0 || j >= 0 || carry != 0 {
		x := carry
		if i >= 0 {
			x += arr1[i]
		}
		if j >= 0 {
			x += arr2[j]
		}

		// 3,2,1,0,-1
		switch {
		case x >= 2:
			// 和正数系统还真不一样，即使表示的是负数位，进位仍然是-1
			ans = append(ans, x-2)
			carry = -1
		case x >= 0:
			ans = append(ans, x)
			carry = 0
		default:
			ans = append(ans, 1)
			carry = 1
		}

		i--
		j--
	}

	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	n := len(ans)
	for k := 0; k < n/2; k++ {
		ans[k], ans[n-1-k] = ans[n-1-k], ans[k]
	}
	return ans
}
