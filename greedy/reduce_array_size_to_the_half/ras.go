package reduce_array_size_to_the_half

import "sort"

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}

	occ := make([]int, 0, len(m))
	for _, c := range m {
		occ = append(occ, c)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(occ)))

	var cnt, ans int
	for _, c := range occ {
		cnt += c
		ans++
		if cnt*2 >= len(arr) {
			break
		}
	}

	return ans
}
