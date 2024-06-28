package valid_mountain_array

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	var i int
	for i = 0; i < n-1 && arr[i+1] > arr[i]; i++ {
	}
	if i >= n-1 || i == 0 {
		return false
	}

	for ; i < n-1 && arr[i+1] < arr[i]; i++ {
	}

	return i == n-1
}
