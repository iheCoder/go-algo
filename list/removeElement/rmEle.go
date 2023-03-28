package removeElement

// 比我想象中要难，确实也有可能是被删除的相互交换
//func removeElement(nums []int, val int) int {
//	n := len(nums)
//	j := n - 1
//	l := 0
//	for i := 0; j < i; i++ {
//		for nums[j] != val {
//			l++
//			j--
//			if j < 0 {
//				return l
//			}
//		}
//		if nums[i] == val {
//			nums[i], nums[j] = nums[j], nums[i]
//		} else {
//			l++
//		}
//	}
//	return l
//}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	j := n - 1
	for nums[j] == val {
		j--
		if j < 0 {
			return j + 1
		}
	}
	for i := 0; j > i; i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			if j <= i {
				return j + 1
			}
			for nums[j] == val {
				j--
			}
		}
	}
	return j + 1
}
