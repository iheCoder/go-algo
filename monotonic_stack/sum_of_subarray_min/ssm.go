package sum_of_subarray_min

func sumSubarrayMins(arr []int) int {
	const mod int = 1e9 + 7
	n := len(arr)

	var ms []int
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for len(ms) > 0 && arr[ms[len(ms)-1]] >= arr[i] {
			ms = ms[:len(ms)-1]
		}

		if len(ms) == 0 {
			l[i] = i + 1
		} else {
			l[i] = i - ms[len(ms)-1]
		}

		ms = append(ms, i)
	}

	ms = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(ms) > 0 && arr[ms[len(ms)-1]] > arr[i] {
			ms = ms[:len(ms)-1]
		}

		if len(ms) == 0 {
			r[i] = n - i
		} else {
			r[i] = ms[len(ms)-1] - i
		}

		ms = append(ms, i)
	}

	var ans int
	for i := 0; i < n; i++ {
		ans = (ans + l[i]*r[i]*arr[i]) % mod
	}
	return ans
}
