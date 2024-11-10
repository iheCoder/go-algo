package decrease_elem_to_make_arr_zigzag

func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	var oddMv, evenMv int
	for i := 0; i < n; i++ {
		var ld, rd int
		if i > 0 && nums[i] >= nums[i-1] {
			ld = nums[i] - nums[i-1] + 1
		}
		if i < n-1 && nums[i] >= nums[i+1] {
			rd = nums[i] - nums[i+1] + 1
		}
		if i%2 == 1 {
			oddMv += maxInt(ld, rd)
		} else {
			evenMv += maxInt(ld, rd)
		}
	}

	return minInt(oddMv, evenMv)
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

// 	if len(nums) == 1 {
//		return 0
//	}
//
//	n := len(nums)
//	dp := make([][2]int, n)
//	numsEven := make([]int, n)
//	copy(numsEven, nums)
//
//	for i := 1; i < n; i++ {
//		if i%2 == 1 {
//			diff := nums[i] - nums[i-1]
//			dp[i][0] = dp[i-1][0]
//			if diff <= 0 {
//				dp[i][0] -= diff - 1
//				nums[i] -= diff - 1
//			}
//
//			diff = numsEven[i] - numsEven[i-1]
//			dp[i][1] = dp[i-1][1]
//			if diff >= 0 {
//				dp[i][1] += diff + 1
//				numsEven[i] += diff + 1
//			}
//		} else {
//			diff := nums[i] - nums[i-1]
//			dp[i][0] = dp[i-1][0]
//			if diff >= 0 {
//				dp[i][0] += diff + 1
//				nums[i] += diff + 1
//			}
//
//			diff = numsEven[i] - numsEven[i-1]
//			dp[i][1] = dp[i-1][1]
//			if diff <= 0 {
//				dp[i][1] -= diff - 1
//				numsEven[i] -= diff - 1
//			}
//		}
//	}
//
//	return minInt(dp[n-1][0], dp[n-1][1])
