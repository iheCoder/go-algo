package next_greater_ele_ii

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{-1}
	}

	stk := make([]int, 0, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	for i := 0; i < 2*n-1; i++ {
		for len(stk) > 0 && nums[stk[len(stk)-1]] < nums[i%n] {
			ans[stk[len(stk)-1]] = nums[i%n]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i%n)
	}
	return ans
}
