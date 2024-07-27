package longest_turbulent_subarray

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	ans := 1
	var left, right int
	for right < n-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right] > arr[right+1] && arr[right-1] < arr[right] {
				right++
			} else if arr[right] < arr[right+1] && arr[right] < arr[right-1] {
				right++
			} else {
				left = right
			}
		}

		ans = maxInt(ans, right-left+1)
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
