package sort_color

func sortColorsMap(nums []int) {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	cnt0 := m[0]
	var i int
	for ; i < cnt0; i++ {
		nums[i] = 0
	}
	cnt1 := m[1]
	for ; i < cnt0+cnt1; i++ {
		nums[i] = 1
	}
	cnt2 := m[2]
	for ; i < cnt0+cnt1+cnt2; i++ {
		nums[i] = 2
	}
}
