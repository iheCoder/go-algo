package peak_index_in_a_mountain

import "sort"

func peakIndexInMountainArray(arr []int) int {
	x := sort.Search(len(arr)-1, func(i int) bool {
		return arr[i+1] <= arr[i]
	})
	//x := sort.Search(len(arr), func(i int) bool {
	//	if i == 0 {
	//		return false
	//	}
	//	return arr[i-1] >= arr[i]
	//})
	return x
}
