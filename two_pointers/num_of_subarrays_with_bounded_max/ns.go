package num_of_subarrays_with_bounded_max

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := func(lower int) int {
		var res int
		c := 0
		for _, num := range nums {
			if num <= lower {
				c++
			} else {
				c = 0
			}
			res += c
		}
		return res
	}

	return count(right) - count(left-1)
}

//func getCombinationCount(n int) int {
//	return (n * (n + 1)) / 2
//}

// 	for {
//		for i < n && nums[i] < left || nums[i] > right {
//			i++
//		}
//		if i >= n {
//			return ans
//		}
//		j = i
//		for j < n && nums[j] >= left && nums[j] <= right {
//			j++
//		}
//		if j == n {
//			return ans + getCombinationCount(j-i)
//		}
//
//		ans += getCombinationCount(j - i + 1)
//		i = j
//	}
//	return ans
