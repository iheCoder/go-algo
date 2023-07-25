package rm_k_digits

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	if k == len(num) {
		return "0"
	}
	arr := make([]int, 0, k)
	// 0大坑会将1、2等等都去除
	push := func(i int) {
		for j := len(arr) - 1; j >= 0 && num[arr[j]] > num[i]; j-- {
			arr = arr[:len(arr)-1]
		}
		arr = append(arr, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]byte, 0, len(num)-k)
	right := -1
	for j := k; j < len(num); j++ {
		push(j)
		for arr[0] <= right {
			arr = arr[1:]
		}
		ans = append(ans, num[arr[0]])
		right = arr[0]
		arr = arr[1:]
	}

	for len(ans) > 0 && ans[0] == '0' {
		ans = ans[1:]
	}
	if len(ans) == 0 {
		return "0"
	}
	return string(ans)
}

// 		if k+i+1 < len(num) && num[k+i+1] < num[j] {
//			j = k + i + 1
//		}

// 	// 若arr都小于k之后窗口的数，那么是不是还是要对k之后进行排序呢
//	for j < len(num) {
//		ans = append(ans, num[j])
//		j++
//	}
