package k_diff_pairs_in_an_array

import "sort"

func findPairs(nums []int, k int) int {
	if len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var target int
	var r int
	var ok bool
	var count int
	max := nums[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target = nums[i] + k
		if target > max {
			return r
		}

		m[nums[i]]--
		if count, ok = m[target]; ok && count > 0 {
			r++
			delete(m, target)
		}
	}

	return r
}
