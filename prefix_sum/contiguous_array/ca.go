package contiguous_array

func findMaxLength(nums []int) int {
	var ans int
	var counter int
	m := map[int]int{0: -1}
	for i, num := range nums {
		if num > 0 {
			counter++
		} else {
			counter--
		}

		if preI, ok := m[counter]; ok {
			ans = maxInt(ans, i-preI)
		} else {
			m[counter] = i
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
