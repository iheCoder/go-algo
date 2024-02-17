package binary_subarrays_with_sum

// 在小于和等于的区间中，都是可有可无的，那么这些可有可无的数量就是便是组合数
func numSubarraysWithSum(nums []int, goal int) int {
	var l, r int
	var ans int
	var ws, l1 int
	for r < len(nums) {
		ws += nums[r]
		// l不能小于等于r，这会导致l比r更大，并且l、r得到的和并不等于ws
		for l < r && ws > goal {
			ws -= nums[l]
			l++
		}
		// 求的是以r为右边界，l为左边界中，l到恰好等于goal之间0的个数
		if ws == goal {
			for l1 = l; l1 < r && nums[l1] == 0; l1++ {
			}
			ans += l1 - l + 1
		}
		r++
	}
	return ans
}
