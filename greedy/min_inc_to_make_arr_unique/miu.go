package min_inc_to_make_arr_unique

import "sort"

// 居然超出时间限制了，难道还是要重新排序吗？
// 确实是要重新排序
func minIncrementForUnique(nums []int) int {
	m := make(map[int]int)
	repeat := make([]int, 0)
	for _, num := range nums {
		m[num]++
		if m[num] == 2 {
			repeat = append(repeat, num)
		}
	}

	if len(repeat) == 0 {
		return 0
	}

	sort.Ints(repeat)
	var ans int
	x := repeat[0] + 1
	for _, num := range repeat {
		if x < num {
			x = num + 1
		}
		for count := m[num]; count > 1; count-- {
			_, ok := m[x]
			for ok {
				x++
				_, ok = m[x]
			}
			m[x] = 1
			ans += x - num
			x++
		}
	}
	return ans
}
