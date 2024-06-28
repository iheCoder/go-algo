package k_th_smallest_prime_fraction

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0., 1.
	for {
		mid := (left + right) / 2
		i, count := -1, 0
		// x,y 代表前一个分子、分母
		x, y := 0, 1

		for j := 1; j < n; j++ {
			// 若当前分数小于中间值，则i向后移动，使的分数更大
			for float64(arr[i+1])/float64(arr[j]) < mid {
				i++
				// 若新分数比之前更大，则更新x、y
				if arr[i]*y > arr[j]*x {
					x, y = arr[i], arr[j]
				}
			}
			//  1/2 j=1 search count+=0
			//	1/3 2/3 j=2 search count+=1
			//	1/5 2/5 3/5 j=3 search count+=2
			// 我理解每轮符合条件的个数一定是i+1个
			// 可为什么最后得到的count一定是第count个呢？
			// 因为这就是比所有的组合
			count += i + 1
		}

		if count == k {
			return []int{x, y}
		}
		if count < k {
			left = mid
		} else {
			right = mid
		}
	}
}
