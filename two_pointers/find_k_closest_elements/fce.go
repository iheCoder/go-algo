package find_k_closest_elements

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	r := make([]int, 0, k)
	i := sort.SearchInts(arr, x)
	if i >= len(arr) {
		i = len(arr) - 1
	}
	if i > 0 && arr[i] != x {
		i--
	}
	j := i + 1
	for m := 0; m < k; m++ {
		if i < 0 || (j < len(arr) && arr[j]-x < x-arr[i]) {
			r = append(r, arr[j])
			j++
		} else {
			r = append(r, arr[i])
			i--
		}
	}

	sort.Ints(r)
	return r
}
