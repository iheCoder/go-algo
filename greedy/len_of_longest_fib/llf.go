package len_of_longest_fib

func lenLongestFibSubseq(arr []int) int {
	m := make(map[int]int)
	for i, num := range arr {
		m[num] = i
	}

	n := len(arr)
	var ans int
	for i := 0; i < n-2; i++ {
		if n-i <= ans {
			break
		}
		for j := i + 1; j < n-1; j++ {
			x, y := i, j
			tempLen := 2
			for {
				a := arr[x] + arr[y]
				k, ok := m[a]
				if !ok {
					break
				}
				x = y
				y = k
				tempLen++
			}

			if tempLen > ans {
				ans = tempLen
			}
		}
	}

	if ans < 3 {
		return 0
	}
	return ans
}
