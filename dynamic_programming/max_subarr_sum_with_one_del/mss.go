package max_subarr_sum_with_one_del

func maximumSum(arr []int) int {
	d0, d1, res := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		d0, d1 = maxInt(d0, 0)+arr[i], maxInt(d0, d1+arr[i])
		res = maxInt(res, maxInt(d0, d1))
	}

	return res
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
