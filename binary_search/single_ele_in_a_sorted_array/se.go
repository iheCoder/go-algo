package single_ele_in_a_sorted_array

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

// 与其计算一半，不如判断mid
// 	if n == 3 {
//		if nums[0] == nums[1] {
//			return nums[2]
//		}
//		return nums[0]
//	}
//	half := n / 2
//	var isEven bool
//	var x int
//	sort.Search(n, func(i int) bool {
//		defer func() {
//			half /= 2
//		}()
//		isEven = half%2 == 0
//		if i > 0 && nums[i] == nums[i-1] {
//			if isEven {
//				return true
//			}
//			half--
//			return false
//		}
//		if i < n-1 && nums[i] == nums[i+1] {
//			if isEven {
//				half--
//				return false
//			}
//			return true
//		}
//		x = i
//		return true
//	})
//
//	return nums[x]
