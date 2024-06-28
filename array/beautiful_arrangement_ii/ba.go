package beautiful_arrangement_ii

func constructArray(n int, k int) []int {
	nums := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		nums = append(nums, i)
	}

	for i, j := n-k, n; i <= j; i++ {
		nums = append(nums, i)
		if i != j {
			nums = append(nums, j)
		}
		j--
	}
	return nums
}
