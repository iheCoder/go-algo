package greatest_sum_divisible_by_3

func maxSumDivThree(nums []int) int {
	maxNum := 1<<32 - 1
	// 1. 找到余1、余2最小的两个数，并统计数量c1、c2与总和
	var ans int
	var p1, p2, c1, c2 int
	n1, n2 := maxNum, maxNum
	for _, num := range nums {
		ans += num
		switch num % 3 {
		case 1:
			c1++
			if num < n1 {
				p1 = n1
				n1 = num
			} else if p1 > num {
				p1 = num
			}
		case 2:
			c2++
			if num < n2 {
				p2 = n2
				n2 = num
			} else if p2 > num {
				p2 = num
			}
		}
	}

	// 2. 若c1等于c2，那么直接返回
	if c1 == c2 {
		return ans
	}
	// 3. 若c1>c2，且c1-c2余1，那么在余2至少有两个情况下，比较余2最小两数之和
	// 若c1-c2余2，在至少有一个余2的情况下，比较最小余2
	if c1 > c2 {
		rem := (c1 - c2) % 3
		if rem == 1 {
			if c2 >= 2 && n1 > n2+p2 {
				ans -= n2 + p2
			} else {
				ans -= n1
			}
		} else if rem == 2 {
			if c2 > 0 && n1+p1 > n2 {
				ans -= n2
			} else {
				ans -= n1 + p1
			}
		}
	} else {
		rem := (c2 - c1) % 3
		if rem == 1 {
			if c1 >= 2 && n2 > n1+p1 {
				ans -= n1 + p1
			} else {
				ans -= n2
			}
		} else if rem == 2 {
			if c1 > 0 && n2+p2 > n1 {
				ans -= n1
			} else {
				ans -= n2 + p2
			}
		}
	}
	// 4. 若c1<c2，且c2-c1余1，那么在余1至少有两个情况下，比较余1最小两数之和
	// 若c2-c1余2，在至少有一个余1的情况下，那么比较最小余1
	return ans
}

// 	sort.Ints(nums)
//
//	var sum1, sum2, last1, count1, count2, last2 int
//	var ans int
//	for i := len(nums) - 1; i >= 0; i-- {
//		switch nums[i] % 3 {
//		case 1:
//			if count1 >= 2 {
//				// 最后两数之和应该减去前面数字而不是后面的数组
//				sum1 -= last1
//			}
//			sum1 += nums[i]
//			last1 = nums[i]
//			count1++
//		case 2:
//			if count2 >= 2 {
//				sum2 -= last2
//			}
//			sum2 += nums[i]
//			last2 = nums[i]
//			count2++
//		}
//		ans += nums[i]
//	}
//
//	if count1 > count2 {
//		rem := (count1 - count2) % 3
//		if rem == 1 {
//			if count2 >= 2 && last1 > sum2 {
//				ans -= sum2
//			} else {
//				ans -= last1
//			}
//		} else if rem == 2 {
//			if count2 > 0 && sum1 > last2 {
//				ans -= last2
//			} else {
//				ans -= sum1
//			}
//		}
//	} else if count2 > count1 {
//		rem := (count2 - count1) % 3
//		if rem == 1 {
//			if count1 >= 2 && last2 > sum1 {
//				ans -= sum1
//			} else {
//				ans -= last2
//			}
//		} else if rem == 2 {
//			if count1 > 0 && sum2 > last1 {
//				ans -= last1
//			} else {
//				ans -= sum2
//			}
//		}
//	}
//	return ans
