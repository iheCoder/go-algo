package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := make([]int, 0)
	for x > 0 {
		arr = append(arr, x%10)
		x /= 10
	}
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] != arr[n-1-i] {
			return false
		}
	}
	return true
}

//func pow10(x int) int {
//	return int(math.Pow10(x))
//}
