package _sum

import "sort"

// 我想通过比较start和最大的三个之和sum，若sum小于target肯定i是不对的。唯一的疑问是

func fourSumDep(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	r := make([][]int, 0)

	end := len(nums) - 1

	var start int
	var i, j int
	for end >= 3 {
		start = 0
		for end-start >= 3 {
			i = start + 1
			j = end - 1
			for i < j {
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					r = append(r, []int{nums[start], nums[i], nums[j], nums[end]})
					i++
					j--
				} else if sum > target {
					j--
				} else {
					i++
				}
			}
			start++
		}
		end--
	}

	return r
}
