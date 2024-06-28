package smallest_range_ii

import "sort"

func smallestRangeII(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 0; i < n-1; i++ {
		a, b := nums[i], nums[i+1]
		hi, lo := maxInt(nums[n-1]-k, a+k), minInt(nums[0]+k, b-k)
		ans = minInt(ans, hi-lo)
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//func smallestRangeII(nums []int, k int) int {
//	if len(nums) == 1 {
//		return 0
//	}
//
//	sort.Ints(nums)
//	n := len(nums)
//
//	diff := nums[n-1] - nums[0]
//	if diff <= k {
//		return diff
//	}
//
//	right, left := nums[n-1]-k, nums[0]+k
//	if left > right {
//		left, right = right, left
//	}
//
//	for i := 1; i < n-1; i++ {
//		ir, il := nums[i]+k, nums[i]-k
//		if ir <= right || il >= left {
//			continue
//		}
//
//		rdiff, ldiff := ir-left, right-il
//		if rdiff <= ldiff {
//			right = ir
//		} else {
//			left = il
//		}
//	}
//
//	d := right - left
//	if diff < d {
//		return diff
//	}
//	return d
//}
