package next_permutation

// 有一个思路。两个指针从后往前，交换前面的值恰好小于后面的值
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	} else if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			j := len(nums) - 1
			for j >= 0 && nums[i] >= nums[j] {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]

			j, k := i+1, len(nums)-1
			for j < k {
				nums[j], nums[k] = nums[k], nums[j]
				j++
				k--
			}

			//k = sort.SearchInts(nums[i+1:], nums[i]+1)
			//nums[i], nums[k+i+1] = nums[k+i+1], nums[i]
			return
		}
	}

	j, k := 0, len(nums)-1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}

//func nextPermutation(nums []int) {
//	n := len(nums)
//	i := n - 2
//	for i >= 0 && nums[i] >= nums[i+1] {
//		i--
//	}
//	if i >= 0 {
//		j := n - 1
//		for j >= 0 && nums[i] >= nums[j] {
//			j--
//		}
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//	reverse(nums[i+1:])
//}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 	last := len(nums) - 1
//	start := last
//	for i := len(nums) - 2; i >= 0; i-- {
//		if nums[i] > nums[last] {
//			nums = append(nums, nums[i])
//			last++
//		} else if nums[i] < nums[last] {
//			k := sort.SearchInts(nums[start:], nums[i])
//			nums[i], nums[k+start] = nums[k+start], nums[i]
//			copy(nums[i+1:start], nums[start:])
//			return
//		}
//	}
//
//	j, k := 0, start
//	for j < k {
//		nums[j], nums[k] = nums[k], nums[j]
//		j++
//		k--
//	}
