package sum_of_mutated_array_closet_to_target

import "sort"

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + arr[i]
	}

	l, r, ans := 0, arr[n-1], -1
	for l <= r {
		mid := (l + r) / 2
		index := sort.SearchInts(arr, mid)
		if index < 0 {
			index = -index - 1
		}

		cur := prefix[index] + (n-index)*mid
		if cur <= target {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	chooseSmall := check(arr, ans)
	chooseBig := check(arr, ans+1)
	if abs(chooseSmall-target) > abs(chooseBig-target) {
		ans++
	}

	return ans
}

func check(arr []int, x int) int {
	var ret int
	for _, num := range arr {
		ret += minInt(num, x)
	}

	return ret
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func findBestValueOneBS(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + arr[i]
	}

	diff := target
	var ans int
	for num := 1; num <= arr[n-1]; num++ {
		i := sort.SearchInts(arr, num)
		if i < 0 {
			i = -i - 1
		}

		cur := prefixSum[i] + (n-i)*num
		if abs(cur-target) < diff {
			ans = num
			diff = abs(cur - target)
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
