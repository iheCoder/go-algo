package min_time_diff

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	m := make(map[int]bool)
	var key int
	for _, point := range timePoints {
		key = getKey(point)
		if m[key] {
			return 0
		}
		m[key] = true
	}

	n := len(m)
	arr := make([]int, 0, n)
	for tp, _ := range m {
		arr = append(arr, tp)
	}

	sort.Ints(arr)
	arr = append(arr, arr[0]+24*60)
	ans := math.MaxInt
	var diff int
	for i := 1; i <= n; i++ {
		diff = arr[i] - arr[i-1]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

func getKey(point string) int {
	ts := strings.Split(point, ":")
	h, _ := strconv.Atoi(ts[0])
	m, _ := strconv.Atoi(ts[1])
	return h*60 + m
}
