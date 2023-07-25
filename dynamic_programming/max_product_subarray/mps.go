package max_product_subarray

// 如果定义dp[i]为以i元素结尾的最大积，那么dp[i]、dp[i-1]关系是什么呢？
// 当nums[i]为正数，nums[i-1]也正数，自然dp[i]=nums[i]*dp[i-1]
// 当nums[i]为负数，nums[i-1]也负数，那么一直乘，直到积为负数为止
// 只要存在负数，就基本上要一直往前找看是否能让负数积变为正
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	ans := -1 << 31
	ps := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(ps); j++ {
			ps[j] *= nums[i]
			ans = maxInt(ans, ps[j])
		}
		ps = append(ps, nums[i])
		ans = maxInt(ans, nums[i])
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 	dp := make([]int, len(nums))
//	dp[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i-1]>=0 {
//			if nums[i]>=0 {
//				dp[i]=maxInt(dp[i-1]*nums[i],nums[i])
//			}else {
//
//			}
//		}else {
//
//		}
//	}
