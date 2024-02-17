package find_peak_ele

func findPeakElement(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	f := func(i int) int {
		lup, rup := true, false
		if i > 0 && nums[i-1] > nums[i] {
			lup = false
		}
		if i < n-1 && nums[i+1] > nums[i] {
			rup = true
		}

		if lup && !rup {
			return 0
		} else if !lup && !rup {
			return -1
		} else {
			return 1
		}
	}
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		switch f(mid) {
		case -1:
			high = mid
		case 1:
			low = mid + 1
		default:
			return mid
		}
	}
	return mid
}
