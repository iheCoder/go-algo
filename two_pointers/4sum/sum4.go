package _sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	var i, j, k, l int
	var sum, totalSum int
	r := make([][]int, 0)
	for i = 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j = i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum = nums[i] + nums[j]
			k = j + 1
			l = len(nums) - 1
			for k < l {
				if k > j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}
				if l < len(nums)-1 && nums[l] == nums[l+1] {
					l--
					continue
				}

				totalSum = sum + nums[k] + nums[l]
				if totalSum > target {
					l--
				} else if totalSum < target {
					k++
				} else {
					r = append(r, []int{nums[i], nums[j], nums[k], nums[l]})
					l--
					k++
				}
			}
		}
	}

	return r
}
