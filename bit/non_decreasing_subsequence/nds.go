package non_decreasing_subsequence

import "math"

var (
	temp []int
	ans  [][]int
)

func findSubsequences(nums []int) [][]int {
	if len(nums) < 2 {
		return nil
	}

	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

func dfs(cur, last int, nums []int) {
	if cur >= len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}

	if last <= nums[cur] {
		temp = append(temp, nums[cur])
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	// 仅当当前元素不等于上一个元素的时候才不选择当前元素
	if last != nums[cur] {
		dfs(cur+1, last, nums)
	}
}

// 	var r [][]int
//	var start int
//	for i := 1; start < len(nums); i++ {
//		if i >= len(nums) || nums[i] < nums[i-1] {
//			r = append(r, getContinuousSubsets(nums[start:i])...)
//			start = i
//		}
//	}
//	return r

// 这确实是获取连续的从2到n的子集，然后却并不真的需要连续
func getContinuousSubsets(arr []int) [][]int {
	if len(arr) < 2 {
		return nil
	}
	n := len(arr)
	var r [][]int
	for c := 2; c <= n; c++ {
		for mask := 1<<c - 1; mask < 1<<n; mask <<= 1 {
			ri := make([]int, 0, c)
			for i := 0; i < n; i++ {
				if mask>>i&1 > 0 {
					ri = append(ri, arr[i])
				}
			}
			r = append(r, ri)
		}
	}
	return r
}
