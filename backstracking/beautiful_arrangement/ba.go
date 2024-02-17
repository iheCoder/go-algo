package beautiful_arrangement

func countArrangement(n int) int {
	um := make([]int, n)
	var ans int
	var f func(i int)
	f = func(i int) {
		if i <= 0 {
			ans++
			return
		}

		for j, v := range um {
			if v > 0 {
				continue
			}
			if (j+1)%i != 0 && i%(j+1) != 0 {
				continue
			}
			um[j] = 1
			f(i - 1)
			um[j] = 0
		}
	}
	f(n)
	return ans
}
