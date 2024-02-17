package pancake_sorting

func pancakeSort(arr []int) []int {
	if len(arr) <= 1 {
		return []int{}
	}

	m := make(map[int]int)
	for i, num := range arr {
		m[num] = i
	}

	r := make([]int, 0)
	var j, k int
	for i := len(arr) - 1; i > 0; i-- {
		k = m[i+1]
		if k == i {
			continue
		}
		if k > 0 {
			r = append(r, k+1)
			j = 0
			for j < k {
				m[arr[j]] = k
				m[arr[k]] = j
				arr[j], arr[k] = arr[k], arr[j]
				j++
				k--
			}
		}
		j = 0
		k = i
		for j < k {
			m[arr[j]] = k
			m[arr[k]] = j
			arr[j], arr[k] = arr[k], arr[j]
			j++
			k--
		}
		r = append(r, i+1)
	}

	return r
}

func reverseArr(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
