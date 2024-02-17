package max_swap

import "sort"

func maximumSwap(num int) int {
	arr := make([]int, 0, 8)
	m := make(map[int]int)
	var rem int
	for i := 0; num > 0; i++ {
		rem = num % 10
		if _, ok := m[rem]; !ok {
			m[rem] = i
		}
		arr = append(arr, rem)
		num /= 10
	}

	n := len(arr)
	sa := make([]int, n)
	copy(sa, arr)
	sort.Ints(sa)

	for i := n - 1; i >= 0; i-- {
		if arr[i] < sa[i] {
			j := m[sa[i]]
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}

	var ans int
	for i := n - 1; i >= 0; i-- {
		ans = ans*10 + arr[i]
	}
	return ans
}
