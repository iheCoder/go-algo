package max_chunks_to_make_sorted

func maxChunksToSorted(arr []int) int {
	edge, expected := 0, 0
	n := len(arr)
	m := make([]int, n)
	var ans int
	for i := 0; i < n; i++ {
		m[arr[i]]++
		if arr[i] > edge {
			edge = arr[i]
		}
		if expected == arr[i] {
			if edge == i {
				ans++
			}
			for j := expected + 1; j < n; j++ {
				if m[j] == 0 {
					expected = j
					break
				}
			}
		}
	}

	return ans
}
