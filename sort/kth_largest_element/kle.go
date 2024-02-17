package kth_largest_element

// 不理解为什么会超时？
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	k = n - k
	// l定义为第几小
	var f func(l, h int) int
	f = func(l, h int) int {
		// 如果到达递归到最后一个，说明l要么是k，要么在第k位之前
		// 对于前者直接返回nums[k]是可以的
		// 对于后者，能到最后一个说明k已经被排序好了
		if l == h {
			return nums[k]
		}

		x := nums[l]
		i, j := l-1, h+1
		for i < j {
			i++
			j--
			for i < len(nums) && nums[i] < x {
				i++
			}
			for j >= 0 && nums[j] > x {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		// 可能i会一直不变，如果
		// j是恰好不小于的分界线
		// 也就是x的第几位
		if k <= j {
			return f(l, j)
		} else {
			return f(j+1, h)
		}
	}

	return f(0, n-1)
}

// 		if nums[i] > x {
//			i--
//		} else {
//			nums[i], nums[l] = nums[l], nums[i]
//			j++
//		}
// 这样会导致超出界限
