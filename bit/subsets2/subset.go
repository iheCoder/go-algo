package subsets2

import "sort"

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	n := len(nums)

	var flag bool
	var ans [][]int
	for mask := 0; mask < (1 << n); mask++ {
		flag = true
		var arr []int
		for i := 0; i < n; i++ {
			if 1&(mask>>i) > 0 {
				if i-1 >= 0 && 1&(mask>>(i-1)) == 0 && nums[i-1] == nums[i] {
					flag = false
					break
				}
				arr = append(arr, nums[i])
			}
		}
		if flag {
			ans = append(ans, arr)
		}
	}
	return ans
}
