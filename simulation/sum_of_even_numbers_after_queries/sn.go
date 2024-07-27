package sum_of_even_numbers_after_queries

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	var sum int
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}

	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		v := query[0]
		i := query[1]

		old := nums[i]
		x := old + v
		if old%2 == 1 || old%2 == -1 {
			if x%2 == 0 {
				sum += x
			}
		} else {
			if x%2 == 0 {
				sum += x - old
			} else {
				sum -= old
			}
		}

		nums[i] = x
		ans = append(ans, sum)
	}
	return ans
}
