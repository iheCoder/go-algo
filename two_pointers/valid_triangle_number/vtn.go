package valid_triangle_number

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	var r int
	var sum int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			sum = nums[i] + nums[j]
			for k := j + 1; k < len(nums); k++ {
				if sum <= nums[k] {
					break
				}
				r++
			}
		}
	}

	return r
}

func isTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		// 任意边长小于等于0，不满足三边不能同时为0的条件
		return false
	}
	if a+b > c && a+c > b && b+c > a {
		// 任意两边之和大于第三边，且任意两边之差小于第三边
		return true
	}
	return false
}
