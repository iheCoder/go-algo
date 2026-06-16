package max_num_of_non_overlapping_subarr_with_sum_equals_target

func maxNonOverlapping(nums []int, target int) int {
	set := map[int]bool{0: true}
	pre, ans := 0, 0

	for _, num := range nums {
		pre += num
		if set[pre-target] {
			ans++
			set = map[int]bool{0: true}
			pre = 0
		} else {
			set[pre] = true
		}
	}

	return ans
}
