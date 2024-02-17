package first_bad_version

import "sort"

func firstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool {
		return isBadVersion(version)
	})
}

func firstBadVersionBS(n int) int {
	if isBadVersion(1) {
		return 1
	}

	low, high := 1, n
	mid := low + (high-low)/2
	vals := make([]int, n+1)
	for {
		if isBadVersion(mid) {
			if mid > 0 && vals[mid-1] == 2 {
				return mid
			}
			high = mid - 1
			vals[mid] = 1
		} else {
			if mid < n && vals[mid+1] == 1 {
				return mid + 1
			}
			low = mid + 1
			vals[mid] = 2
		}
		mid = low + (high-low)/2
	}
}

func isBadVersion(version int) bool {
	return false
}

// func firstBadVersion(n int) int {
//	low, high := 1, n
//	mid := low + (high-low)/2
//	p := mid
//	vals := make([]int, n+1)
//	for {
//		if isBadVersion(p) {
//			if p < n && vals[p+1] == 2 {
//				return p + 1
//			}
//			p = low
//			high = mid - 1
//			vals[p] = 1
//		} else {
//			if p > 0 && vals[p-1] == 1 {
//				return p
//			}
//			p = high
//			low = mid + 1
//			vals[p] = 2
//		}
//		mid = low + (high-low)/2
//	}
//}
