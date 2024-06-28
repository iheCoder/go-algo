package non_decreasing_array

func checkPossibility(nums []int) bool {
	n := len(nums)
	var used bool
	for i := 0; i < n-1; i++ {
		if nums[i+1] >= nums[i] {
			continue
		}

		if used {
			return false
		}

		// 如果是i-1个最大了，那么完全是可以忽略i-1个的，这并不影响后续的比较
		// 如果是i最大了，那么在i+1值大于i-1的时候，那么i值可以设置为i+1值，但由于下一次肯定不会再取到该i值，因为如果取到的话，说明第二次出现非递增的，就会return false
		// 如果是i+1值小于i-1时候，那就直接设置i+1值为i-1
		if i > 0 && nums[i+1] < nums[i-1] {
			nums[i+1] = nums[i]
		}

		used = true
	}

	return true
}

// 	nums = append(nums, math.MaxInt)
//	nums = append([]int{math.MinInt}, nums...)
//	for i := 1; i <= n; i++ {
//		if nums[i] >= nums[i-1] && nums[i+1] >= nums[i] {
//			continue
//		}
//		if nums[i] < nums[i-1] && nums[i+1] < nums[i] {
//			return false
//		}
//		if nums[i] >= nums[i-1] && nums[i+1] <= nums[i] {
//			if used {
//				return false
//			}
//			if nums[i+1] >= nums[i-1] {
//				nums[i]=nums[i+1]
//			}else {
//				nums[i+1]=nums[i]
//			}
//			used=true
//		}
//		if nums[i] <= nums[i-1] && nums[i] >= nums[i+1] {
//			if used {
//				return false
//			}
//			if nums[i-1] <= nums[i+1] {
//				nums[i]=nums[i+1]
//			}else {
//
//			}
//		}
//
//	}
