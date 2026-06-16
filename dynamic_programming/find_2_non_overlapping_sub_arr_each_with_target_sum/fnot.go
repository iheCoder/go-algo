package find_2_non_overlapping_sub_arr_each_with_target_sum

// You are given an array of integers arr and an integer target.
//
//You have to find two non-overlapping sub-arrays of arr each with a sum equal target. There can be multiple answers so you have to find an answer where the sum of the lengths of the two sub-arrays is minimum.
//
//Return the minimum sum of the lengths of the two required sub-arrays, or return -1 if you cannot find such two sub-arrays.
//
//
//
//Example 1:
//
//Input: arr = [3,2,2,4,3], target = 3
//Output: 2
//Explanation: Only two sub-arrays have sum = 3 ([3] and [3]). The sum of their lengths is 2.
//Example 2:
//
//Input: arr = [7,3,4,7], target = 7
//Output: 2
//Explanation: Although we have three non-overlapping sub-arrays of sum = 7 ([7], [3,4] and [7]), but we will choose the first and third sub-arrays as the sum of their lengths is 2.
//Example 3:
//
//Input: arr = [4,3,2,6,2,3,4], target = 6
//Output: -1
//Explanation: We have only one sub-array of sum = 6.
//
//
//Constraints:
//
//1 <= arr.length <= 105
//1 <= arr[i] <= 1000
//1 <= target <= 108

func minSumOfLengths(arr []int, target int) int {
	// 职责：在线性扫描中同时完成两件事：
	// 1. 用滑动窗口找出所有和为 target 的子数组；
	// 2. 对每个窗口，只和它左侧已经出现过的最短合法窗口拼接。
	//
	// 这里依赖题目约束 arr[i] > 0：窗口和过大时左端点右移后，sum 只会变小，
	// 因而每个元素最多进出窗口一次，不需要前缀和哈希表处理负数场景。
	n := len(arr)
	impLen := (n + 1) * 2
	impossibleAnswer := impLen
	answer := impossibleAnswer

	bestBef := make([]int, n)
	for i := range bestBef {
		bestBef[i] = impLen
	}
	left, sum := 0, 0
	bestLen := impLen

	for right, val := range arr {
		sum += val
		for sum > target && left < right {
			sum -= arr[left]
			left++
		}

		if sum == target {
			curLen := right - left + 1
			if left > 0 && bestBef[left-1] != impLen {
				answer = minInt(answer, curLen+bestBef[left-1])
			}

			bestLen = minInt(bestLen, curLen)
		}

		bestBef[right] = bestLen
	}

	if answer == impossibleAnswer {
		return -1
	}
	return answer
}

// minInt 保持比较语义集中，避免主流程里反复出现分支噪音。
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
