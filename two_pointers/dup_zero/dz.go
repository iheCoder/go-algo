package dup_zero

func duplicateZeros(arr []int) {
	zeroCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroCount++
		}
	}

	j := len(arr) + zeroCount - 1
	i := len(arr) - 1
	for ; i >= 0 && j >= len(arr); i-- {
		if arr[i] == 0 {
			j--
			if j < len(arr) {
				arr[j] = 0
				i--
				j--
				break
			}
		}
		j--
	}

	for ; i >= 0; i-- {
		if arr[i] != 0 {
			arr[j] = arr[i]
		} else {
			arr[j] = 0
			j--
			arr[j] = 0
		}
		j--
	}
}
