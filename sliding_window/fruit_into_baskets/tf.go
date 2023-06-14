package fruit_into_baskets

func totalFruit(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}

	m := make(map[int]int)
	var ans, i int
	for j := 0; j < len(fruits); j++ {
		if _, ok := m[fruits[j]]; ok || len(m) < 2 {
			ans = maxInt(ans, j-i+1)
		} else {
			i = popPeerItem(m, fruits[j-1]) + 1
		}
		m[fruits[j]] = j
	}
	return ans
}

func popPeerItem(m map[int]int, x int) int {
	for k, v := range m {
		if k != x {
			delete(m, k)
			return v
		}
	}
	return -1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
