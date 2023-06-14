package subarray_product_less_than_k

func numSubarrayProductLessThanK(nums []int, k int) int {
	// 滑动窗口。以子串积作为窗口元素，i、j为左右边界
	// 若nums[x]小于k，那么添加到窗口，并且寻找k/nums[x]作为新的左边界，并且统计此时窗口元素个数
	window := make([]int, 0, len(nums))
	var ans int
	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] >= k {
			i = len(window)
			continue
		}
		ans++
		for x := i; x < len(window); x++ {
			window[x] *= nums[j]
			if window[x] >= k {
				i = x + 1
				continue
			}
			ans++
		}
		window = append(window, nums[j])
	}
	return ans
}

func numSubarrayProductLessThanKPrefixProduct(nums []int, k int) int {
	// 1. 初始化前缀积ps
	ps := make([]int, len(nums)+1)
	ps[0] = 1
	totalProduct := 1
	for i, num := range nums {
		totalProduct *= num
		ps[i+1] = totalProduct
	}

	// 2. 遍历元素i，若i大于等于k跳过；否则遍历ps[i]/ps[i-steps]，知道大于等于k为止
	var ans, p int
	for i := 0; i < len(nums); i++ {
		p = nums[i]
		j := i
		for p < k {
			ans++
			j--
			if j < 0 {
				break
			}
			p = ps[i+1] / ps[j]
		}
	}
	return ans
}

// 前缀积是否可行呢？
// 似乎可行的
// 并不行，积超出了整数能表示的最大范围
