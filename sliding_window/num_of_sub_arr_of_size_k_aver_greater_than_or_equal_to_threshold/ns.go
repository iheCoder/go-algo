package num_of_sub_arr_of_size_k_aver_greater_than_or_equal_to_threshold

func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	var count, winSum int

	for i := 0; i < k; i++ {
		winSum += arr[i]
	}

	if winSum >= threshold*k {
		count++
	}

	for i := k; i < n; i++ {
		winSum += arr[i] - arr[i-k]
		if winSum >= threshold*k {
			count++
		}
	}

	return count
}
