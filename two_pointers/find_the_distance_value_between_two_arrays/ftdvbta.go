package find_the_distance_value_between_two_arrays

import (
	"math"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)

	var count int
	for _, a := range arr1 {
		x := shortestAbsPathBS(a, arr2)
		if x > d {
			count++
		}
	}
	return count
}

func absInt(x int) int {
	return int(math.Abs(float64(x)))
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func shortestAbsPathBS(x int, arr []int) int {
	if x >= arr[len(arr)-1] {
		return absInt(x - arr[len(arr)-1])
	}
	if x <= arr[0] {
		return absInt(x - arr[0])
	}

	var mid int
	l, r := 0, len(arr)-1
	for l < r {
		mid = l + (r-l)/2
		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if r == 0 {
		return absInt(x - arr[r])
	}
	return minInt(absInt(x-arr[r]), absInt(x-arr[r-1]))
}

func binarySearch(x int, arr []int) int {
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-1]
	}
	if x <= arr[0] {
		return arr[0]
	}

	var mid int
	l, r := 0, len(arr)-1
	for l < r {
		mid = l + (r-l)/2
		if arr[mid] > x {
			r = mid - 1
		} else if arr[mid] < x {
			l = mid + 1
		} else {
			return x
		}

	}
	return arr[mid]
}

// 能否想清楚再做啊，明显二分外面还要遍历arr1啊
// 我实在没想到题目解答真是二分查找
// 	mid:=len(arr2)/2
//	l,r:=0,len(arr2)-1
//	for l <= r {
//		if arr2[mid]> {
//
//		}
//	}

// 	var count int
//	for _, a := range arr1 {
//		if absInt(a-arr2[0]) <= d || absInt(a-arr2[len(arr2)-1]) <= d {
//			count++
//		}
//	}
//
//	return len(arr1) - count
