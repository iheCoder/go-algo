package max_width_ramp

func maxWidthRamp(nums []int) int {
	n := len(nums)
	l := n
	var i, j int
	for l > 0 {
		i, j = 0, l-1
		for j < n {
			if nums[j] >= nums[i] {
				return j - i
			}
			i++
			j++
		}
		l--
	}
	return 0
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 暴力解法超时了
// 	for i := 0; i < n-1; i++ {
//		for j := n - 1; j > i; j-- {
//			if nums[i] <= nums[j] {
//				ans = maxInt(ans, j-i)
//			}
//		}
//	}
//	return ans

// 不理解和单调栈有什么关系，明明和中间那些元素没有关系啊
// 	stk := make([]int, 0)
//	var ans int
//	var push func(x int)
//	push = func(x int) {
//		if len(stk) > 0 && x > stk[len(stk)-1] {
//			stk = stk[:len(stk)-1]
//		}
//		stk = append(stk, x)
//		if len(stk) > ans {
//			ans = len(stk)
//		}
//	}
//	for _, num := range nums {
//		push(num)
//	}
//	return ans
