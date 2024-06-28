package split_array_into_consecutive_subsequences

func isPossible(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	left := make(map[int]int)
	for _, num := range nums {
		left[num]++
	}

	endCnt := make(map[int]int)
	for _, num := range nums {
		if left[num] == 0 {
			continue
		}

		if endCnt[num-1] > 0 {
			endCnt[num]++
			left[num]--
			endCnt[num-1]--
		} else if left[num+1] > 0 && left[num+2] > 0 {
			endCnt[num+2]++
			left[num]--
			left[num+1]--
			left[num+2]--
		} else {
			return false
		}
	}

	return true
}
