package array_of_doubled_pairs

import "sort"

// 如果1/2和2倍的都有，且数量相等，或者数量都大于当前值，那么选哪一个？
// 排序啊！
func canReorderDoubled(arr []int) bool {
	pos := make([]int, 0, len(arr)/2)
	neg := make([]int, 0, len(arr)/2)
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			neg = append(neg, arr[i])
		} else {
			pos = append(pos, arr[i])
		}
	}

	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	sort.Ints(pos)
	sort.Sort(sort.Reverse(sort.IntSlice(neg)))

	for _, v := range pos {
		if m[v] <= 0 {
			continue
		}

		if m[v*2] <= 0 {
			return false
		}

		m[v*2]--
		m[v]--
	}

	for _, v := range neg {
		if m[v] <= 0 {
			continue
		}

		if m[v*2] <= 0 {
			return false
		}

		m[v*2]--
		m[v]--
	}

	return true
}
