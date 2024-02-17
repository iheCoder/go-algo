package sliding_window_max

func maxSlidingWindow(nums []int, k int) []int {
	indexQ := make([]int, 0)
	push := func(i int) {
		for len(indexQ) > 0 && nums[i] >= nums[indexQ[len(indexQ)-1]] {
			indexQ = indexQ[:len(indexQ)-1]
		}
		indexQ = append(indexQ, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	r := make([]int, 1, len(nums)-k+1)
	r[0] = nums[indexQ[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for indexQ[0] <= i-k {
			indexQ = indexQ[1:]
		}
		r = append(r, nums[indexQ[0]])
	}
	return r
}
