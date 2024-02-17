package premutations

func permute(nums []int) [][]int {
	n := len(nums)
	m := make(map[string]bool)

	var f func(s []byte, i int)
	f = func(s []byte, i int) {
		m[string(s)] = true
		if i >= n {
			return
		}

		for j := i; j < n; j++ {
			s[i], s[j] = s[j], s[i]
			f(s, i+1)
			s[i], s[j] = s[j], s[i]
		}
	}
	var s []byte
	for i := 0; i < n; i++ {
		s = append(s, '0'+byte(i))
	}
	f(s, 0)

	var ans [][]int
	for k, _ := range m {
		a := make([]int, n)
		for j, x := range k {
			a[j] = nums[x-'0']
		}
		ans = append(ans, a)
	}
	return ans
}
