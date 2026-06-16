package max_chunks_to_make_sorted_ii

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	leftMax := make([]int, n)
	rightMin := make([]int, n)

	leftMax[0] = arr[0]
	for i := 1; i < n; i++ {
		leftMax[i] = maxInt(leftMax[i-1], arr[i])
	}

	rightMin[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMin[i] = minInt(rightMin[i+1], arr[i])
	}

	var ans int
	for i := 0; i < n-1; i++ {
		if leftMax[i] <= rightMin[i+1] {
			ans++
		}
	}

	return ans + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
